package main

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func main() {
	session, err := mgo.Dial("")
	session.SetMode(1,true)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success.")
	}
	defer session.Close()

	db := session.DB("users")
	collection:=db.C("maijia")

	var users []maijia
	
	count,_:=collection.Count()
	for i:=0;i<(count+1)/2;i++{
		collection.Find(nil).Limit(2).Skip(i*2).All(&users)
		fmt.Println(users)
	}
	collection.Find(bson.M{"name": bson.M{"$in": []string{"fouth_auto_mj", "first_auto_mj"}}}).All(&users)
	fmt.Println(users)
}