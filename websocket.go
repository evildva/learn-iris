package main
/*
import (
    "context"
    "log"
    "net/http"
    "os"

    "github.com/kataras/neffos"
    "github.com/kataras/neffos/gorilla"
)

var events = neffos.Namespaces{
    "v1": neffos.Events{
        "echo": onEcho,
    },
}

func onEcho(c *neffos.NSConn, msg neffos.Message) error {
    body := string(msg.Body)
    log.Println(body)

    if !c.Conn.IsClient() {
        newBody := append([]byte("echo back: "), msg.Body...)
        return neffos.Reply(newBody)
    }

    return nil
}

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        log.Fatalf("expected program to start with 'server' or 'client' argument")
    }
    side := args[0]

    switch side {
    case "server":
        runServer()
    case "client":
        runClient()
    default:
        log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
    }
}

func runServer() {
    websocketServer := neffos.New(gorilla.DefaultUpgrader, events)

    router := http.NewServeMux()
    router.Handle("/echo", websocketServer)

    log.Println("Serving websockets on localhost:8080/echo")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func runClient() {
    ctx := context.Background()
    client, err := neffos.Dial(ctx,gorilla.DefaultDialer,"ws://localhost:8080/echo",events)
    if err != nil {
        panic(err)
    }

    c, err := client.Connect(ctx, "v1")
    if err != nil {
        panic(err)
    }

    c.Emit("echo", []byte("Greetings!"))

    // a channel that blocks until client is terminated,
    // i.e by CTRL/CMD +C.
    <-client.NotifyClose
}
*/
import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gobwas"
	"github.com/kataras/neffos/gorilla"
)

/*
	$ go run main.go server
	# new tab(s) and:
	$ go run main.go client
*/

const (
	endpoint  = "localhost:9090"
	namespace = "default"
	timeout   = 0 // 30 * time.Second
)

var handler = neffos.WithTimeout{
	ReadTimeout:  timeout,
	WriteTimeout: timeout,
	Namespaces: neffos.Namespaces{
		"default": neffos.Events{
			neffos.OnNamespaceConnect: func(c *neffos.NSConn, msg neffos.Message) error {
				if msg.Err != nil {
					log.Printf("This client can't connect because of: %v", msg.Err)
					return nil
				}

				err := fmt.Errorf("Server says that you are not allowed here")
				/* comment this to see that the server-side will
				no allow to for this socket to be connected to the "default" namespace
				and an error will be logged to the client. */
				err = nil

				return err
			},
			neffos.OnNamespaceConnected: func(c *neffos.NSConn, msg neffos.Message) error {
				if !c.Conn.IsClient() {
					c.Emit("chat", []byte("welcome to server's namespace"))
				}

				log.Printf("[%s] connected to [%s].", c.Conn.ID(), msg.Namespace)

				return nil
			},
			neffos.OnNamespaceDisconnect: func(c *neffos.NSConn, msg neffos.Message) error {
				if msg.Err != nil {
					log.Printf("This client can't disconnect yet, server does not allow that action, reason: %v", msg.Err)
					return nil
				}

				err := fmt.Errorf("Server says that you are not allowed to be disconnected yet")
				/* here if you comment this, the return error will mean that
				the disconnect message from client-side will be ignored from the server
				and the connection would be still available to send message to the "default" namespace
				it will not be disconnected.*/
				err = nil

				if err == nil {
					log.Printf("[%s] disconnected from [%s].", c.Conn.ID(), msg.Namespace)
				}

				if c.Conn.IsClient() {
					os.Exit(0)
				}

				return err
			},
			"chat": func(c *neffos.NSConn, msg neffos.Message) error {
				if !c.Conn.IsClient() {
					// this is possible too:
					// if bytes.Equal(msg.Body, []byte("force disconnect")) {
					// 	println("force disconnect")
					// 	return c.Disconnect()
					// }

					log.Printf("--server-side-- send back the message [%s:%s]", msg.Event, string(msg.Body))
					//	c.Emit(msg.Event, msg.Body)
					//	c.Server().Broadcast(nil, msg) // to all including this connection.
					c.Conn.Server().Broadcast(c.Conn, msg) // to all except this connection.
				}

				log.Printf("---------------------\n[%s] %s", c.Conn.ID(), msg.Body)
				return nil
			},
		},
	},
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("expected program to start with 'server' or 'client' argument")
	}
	side := args[0]

	var (
		upgrader = gobwas.DefaultUpgrader
		dialer   = gobwas.DefaultDialer
	)

	if len(args) > 1 {
		method := args[1]
		if method == "gorilla" {
			upgrader = gorilla.DefaultUpgrader
			dialer = gorilla.DefaultDialer
			if side == "server" {
				log.Printf("Using with Gorilla Upgrader.")
			} else {
				log.Printf("Using with Gorilla Dialer.")
			}
		}
	}

	switch side {
	case "server":
		server(upgrader)
	case "client":
		client(dialer)
	default:
		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
	}
}

var (
	// tests immediately closed on the `Server#OnConnect`.
	dissalowAll = false
	// if not empty, tests broadcast on `Server#OnConnect` (expect this conn because it is not yet connected to any namespace locally).
	notifyOthers                  = true
	serverHandlesConnectNamespace = true
)

func server(upgrader neffos.Upgrader) {
	srv := neffos.New(upgrader, handler)
	// s, err := redis.NewStackExchange(redis.Config{}, "ch")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// srv.UseStackExchange(s)

	srv.OnConnect = func(c *neffos.Conn) error {
		if dissalowAll {
			return fmt.Errorf("you are not allowed to connect here for some reason")
		}

		log.Printf("[%s] connected to server.", c.ID())

		if serverHandlesConnectNamespace {
			ns, err := c.Connect(nil, namespace)
			if err != nil {
				panic(err)
			}

			ns.Emit("chat", []byte("(Force-connected by server)"))
		}

		if notifyOthers {
			c.Server().Broadcast(c, neffos.Message{
				Namespace: namespace,
				Event:     "chat",
				Body:      []byte(fmt.Sprintf("Client [%s] connected too.", c.ID())),
			})

			// c.Server().Broadcast(c, neffos.Message{
			// 	Namespace: namespace,
			// 	Event:     "chat",
			// 	Body:      []byte(fmt.Sprintf("SECOND ONE")),
			// })
		}

		return nil
	}

	srv.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] disconnected from the server.", c.ID())
	}

	srv.OnUpgradeError = func(err error) {
		log.Printf("ERROR: %v", err)
	}

	log.Printf("Listening on: %s\nPress CTRL/CMD+C to interrupt.", endpoint)
	go http.ListenAndServe(endpoint, srv)

	fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Printf("ERROR: %v", scanner.Err())
			return
		}

		text := scanner.Bytes()
		if bytes.Equal(text, []byte("force disconnect")) {
			srv.Do(func(c *neffos.Conn) {
				c.DisconnectAll(nil)
				//	c.Namespace(namespace).Disconnect(nil)
			}, false)
		} else {
			// srv.Do(func(c neffos.Conn) {
			// 	c.Write(namespace, "chat", text)
			// }, false)
			srv.Broadcast(nil, neffos.Message{Namespace: namespace, Event: "chat", Body: text})
		}
		fmt.Fprint(os.Stdout, ">> ")
	}
}

const dialAndConnectTimeout = 5 * time.Second

func client(dialer neffos.Dialer) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(dialAndConnectTimeout))
	defer cancel()

	client, err := neffos.Dial(ctx, dialer, endpoint, handler)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	var c *neffos.NSConn

	if serverHandlesConnectNamespace {
		c, err = client.WaitServerConnect(ctx, namespace)
	} else {
		c, err = client.Connect(ctx, namespace)
	}

	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Printf("ERROR: %v", scanner.Err())
			return
		}

		text := scanner.Bytes()

		if bytes.Equal(text, []byte("exit")) {
			if err := c.Disconnect(nil); err != nil {
				// log.Printf("from server: %v", err)
			}
			continue
		}

		ok := c.Emit("chat", text)
		if !ok {
			break
		}

		fmt.Fprint(os.Stdout, ">> ")
	}
}
