package Handlers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"

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

func (db *webhookdb)update (d webhookobj) {
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

func (db *webhookdb) retPatt(){
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var resault struct{ text string `json:"text"`}
	for range db.webhookCollection{
		//obj := webhookobj{}
		err = session.DB(db.dbName).C(db.webhookCollection).Find(nil).Select(bson.M{"keyid":1}).One(resault)
		if err != nil{
			log.Println(err)
		}

	}
}

func (db *webhookdb) find(keyId string) webhookobj {
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
func (db *webhookdb) findRates(date string) map[string]float64 {
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


func(db *webhookdb) delete(keyId string) int {
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
func (db *webhookdb) addFixer(m Mother){
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

func (db *webhookdb) count() int{
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

func (db *webhookdb) findAll() []webhookobj{
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

func (db *webhookdb) findAllRates() []Mother {
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