package main

import (
    "context"
    "fmt"

    kivik "github.com/go-kivik/kivik/v3"
    _ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
)

func main() {
    client, err := kivik.New("couch", "http://administer:secretpass@localhost:5984/")
    if err != nil {
        panic(err)
    }

    db := client.DB(context.TODO(), "animals")

    doc := map[string]interface{}{
        "_id":      "cow",
        "feet":     4,
        "greeting": "moo",
    }

    rev, err := db.Put(context.TODO(), "animals", doc)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Cow inserted with revision %s\n", rev)
}
