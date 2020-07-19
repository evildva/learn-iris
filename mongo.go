package main

import (
	"context"
	"fmt"
	"log"
	//"sync"
	//"encoding/json"
	"github.com/satori/go.uuid"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type maijia struct{
	ID     int        `bson:"id"`
	Name   string     `bson:"name"`
	Pass   string     `bson:"pass"`
	Phone  string     `bson:"phone"`
	Addr   string     `bson:"addr"`
	Shop   string     `bson:"shop"`
	Order  string     `bson:"order"`
	More   []string   `bson:"more"`
}

type shop struct{
  Username    string
  ShopName    string
  IntroImage  string
  Vegetables  []vegetable
  More        string
}

type vegetable struct{
  Price       int
  Image       string
  Name        string
  More        string
}

type user struct{
  Name        string
  Pass        string
  Phone       string
  Addr        string
  Shop        string
  Order       string
  More        string
}

/*
type A []interface{}

type D []E

type E struct {
    Name  string
    Value interface{}
}

type M map[string]interface{}

 */

func Page(database *mongo.Database,collection string,limit int64,page int64,doc bson.D) (data []map[string]interface{},err error){
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	findoptions:=options.Find()
	if limit > 0 {
		findoptions.SetLimit(limit)
		findoptions.SetSkip(limit * page)
	}
	cur, err := database.Collection(collection).Find(ctx, doc, findoptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &data)
	fmt.Println(len(data))
	return
}

func find(database *mongo.Database,collection string,doc bson.D) bool{
	cur, err := database.Collection(collection).Find(context.TODO(), doc)
	if err != nil {
		return true
	}
	var data []map[string]interface{}
	cur.All(context.Background(), &data)
	if len(data)>0{
		return true
	}
	fmt.Println(data)
	return false
}

func insert(database *mongo.Database,collection string,doc bson.D) bool{
	_,err:=database.Collection(collection).InsertOne(context.TODO(),doc)
	if err != nil {
    	return false
	}

	return true
}

func Update(database *mongo.Database,collection string,filter interface{},update interface{}) bool{
	_,err:=database.Collection(collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
    	return false
	}

	return true
}

func sele(database *mongo.Database,collection string,doc bson.A) (data []map[string]interface{},err error){
	ctx, cannel := context.WithTimeout(context.Background(), time.Minute)
	defer cannel()
	cur, err := database.Collection(collection).Aggregate(ctx, doc)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &data)
	fmt.Println(len(data))
	return
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
/*
	fmt.Println(Page(client.Database("users"),"maijia",2,1))

	var m []map[string]interface{}
	option:=options.Find()
	option.SetLimit(1)
	collection := client.Database("users").Collection("maijia")

	var count,_=collection.EstimatedDocumentCount(context.TODO())
	var i int64
	for i=0;i<count;i++{
		option.SetSkip(i)
		single,_:=collection.Find(context.TODO(),bson.D{{"phone","1234567"}},option)
		single.All(context.TODO(),&m)
		fmt.Println(m)
	}
*/
	//find(client.Database("users"),"maijia","fouth_auto_mj")
	/*
	doc:=bson.D{{"id",0},{"name","seven_auto_mj"},{"pass","abc"},{"phone","1234567"}}
	er:=insert(client.Database("users"),"maijia",doc)
	if er{
		fmt.Println("yes")
	}
	*/
	filter:=bson.D{{"name","seven_auto_mj"}}
	update:=bson.D{{"$set",bson.M{"addr":"a.b.c.r"}}}
	er:=Update(client.Database("users"),"maijia",filter,update)
	if er{
		fmt.Println("yes")
	}
	fmt.Println(uuid.NewV4())
/*
	var syn sync.WaitGroup
	db1:=client.Database("db1").Collection("dbt1")
	db2:=client.Database("db2").Collection("dbt2")

	syn.Add(2)

	go func(){
		for i:=1;i<=20;i++{
			db1.InsertOne(context.TODO(),bson.D{{"para",i},{"a","b"}})
		}
		syn.Done()
	}()

	go func(){
		for i:=1;i<=20;i++{
			db2.InsertOne(context.TODO(),bson.D{{"para",i}})
		}
		syn.Done()
	}()

	syn.Wait()

	sl:=[]interface{}{"fouth_auto_mj"}
	if isExists(client.Database("users"),"maijia",sl){
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	
	*/
	groupStage := bson.D{{"$project", bson.D{{"_id", 0},{"shopname",1}}}}
	fmt.Println(sele(client.Database("shops"),"shop",bson.A{groupStage}))
	//fmt.Println(Page(client.Database("shops"),"shop",1,0,bson.D{{},{"shopname":1,"image":1}}))
}