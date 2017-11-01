package Handlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"

	"net/http"
	"io/ioutil"
	"encoding/json"
)


func(db *webhookdb) Add(s webhookobj) {
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.dbName).C(db.webhookCollection).Insert(s)
	if err != nil {
		fmt.Printf("error in insert: %v", err.Error())
	}

}

func (db *webhookdb)Update (d webhookobj) {
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.dbName).C(db.webhookCollection).Update(bson.M{"webhookurl": d.WebhookURL},d )
	if err != nil {
		fmt.Printf("error: %v",err.Error())
	}

}

func (db *webhookdb) Find(keyId string) webhookobj {
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	resualt := webhookobj{}

	err = session.DB(db.dbName).C(db.webhookCollection).Find(bson.M{"keyid":keyId}).One(&resualt)
	if err != nil{
		log.Println("error with finding", err.Error())
	}
	log.Println("KeyID for find", keyId)
	log.Println("resault keyID", resualt.KeyId)
	return resualt

}
func (db *webhookdb) FindRates(date string) map[string]float64 {
	session,err  := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	log.Println(date)
	resualt:= Mother{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(bson.M{"date": date}).One(&resualt)
	log.Println("lol",resualt)
	if err != nil{
		log.Println("123",err)
	}
	var res map[string]float64 = resualt.Rates
	return res
}


func(db *webhookdb) Delete(keyId string) int {
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.dbName).C(db.webhookCollection).Remove(bson.M{"keyid": keyId})
	if err != nil{
		log.Println("error:  %v", err.Error())
		return 0
	} else{ return 1}
}
func (db *webhookdb) AddFixer(m Mother){
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.dbName).C(db.webhookCollection).Insert(m)
	if err != nil {
		fmt.Printf("error in insert: %v", err.Error())
	}
}

func (db *webhookdb) Count() int{
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	returnValue, err := session.DB(db.dbName).C(db.webhookCollection).Count()
	if err != nil{
		log.Println("error:  %v", err.Error())
		return 0
	} else{ return returnValue}
}

func (db *webhookdb) FindAll() []webhookobj{
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	allHooks := []webhookobj{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(nil).All(&allHooks)
	if err != nil{
		log.Println(err)
	}
	return  allHooks

}

func (db *webhookdb) FindAllRates() []Mother {
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	allFixers := []Mother{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(nil).All(&allFixers)
	if err != nil{
		log.Println(err)
	}
	return  allFixers
}

func FechtAll(){
	log.Println(Database)
	resp, err := http.Get("http://api.fixer.io/latest")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err.Error())
	}
	values := Mother{}
	err = json.Unmarshal(body, &values)
	if err != nil{
		panic(err.Error())
	}
	log.Println(values)
	FixerColl.AddFixer(values)
}