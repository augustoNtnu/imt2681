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


func(db *webhookdb) Add(s webhookobj) int {
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)

	}else {
		err = session.DB(db.dbName).C(db.webhookCollection).Insert(s)
		if err != nil {
			fmt.Printf("error in insert: %v", err.Error())
			status = 0
		}
		defer session.Close()
	}

	return status
}

func (db *webhookdb) Find(keyId string) (webhookobj, int) {
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	resualt := webhookobj{}

	err = session.DB(db.dbName).C(db.webhookCollection).Find(bson.M{"keyid":keyId}).One(&resualt)
	if err != nil{
		status = 0
		log.Println("error with finding", err.Error())
	}
	log.Println("KeyID for find", keyId)
	log.Println("resault keyID", resualt.KeyId)
	return resualt, status

}
func (db *webhookdb) FindRates(date string) (map[string]float64, int) {
	status := 1
	session,err  := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)

	}
	log.Println("date for find rates", date)
	defer session.Close()
	resualt:= Mother{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(bson.M{"date":date}).One(&resualt)
	if err != nil{
		status = 0
		log.Println("1234",err)
	}
	var res map[string]float64 = resualt.Rates
	return res, status
}
func (db *webhookdb) FindFixer(date string) (Mother,int) {
	status:= 1
	session,err  := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)

	}
	defer session.Close()
	log.Println(date)
	resualt:= Mother{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(bson.M{"date":date}).One(&resualt)
	log.Println("FindFixer resault",resualt)
	if err != nil{
		status = 0
		log.Println("123",err)
	}
	return resualt, status
}

func(db *webhookdb) Delete(keyId string) int {
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)

	}
	defer session.Close()

	err = session.DB(db.dbName).C(db.webhookCollection).Remove(bson.M{"keyid":keyId})
	if err != nil{
		log.Println("error:  ", err.Error())
		status = 0

	}
	 return status
}
func (db *webhookdb) AddFixer(m Mother) int{
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}


	err = session.DB(db.dbName).C(db.webhookCollection).Insert(m)
	if err != nil {
		status = 0
		fmt.Println("error in insert: ", err.Error())
	}
	session.Close()
	return status
}

func (db *webhookdb) Count() int{
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	returnValue, err := session.DB(db.dbName).C(db.webhookCollection).Count()
	if err != nil{
		log.Println("error: ", err.Error())
		return 0

	} else{ return returnValue}
}

func (db *webhookdb) FindAll() ([]webhookobj, int){
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		status = 0
		panic(err)
	}
	defer session.Close()
	allHooks := []webhookobj{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(nil).All(&allHooks)
	if err != nil{
		status = 0
		log.Println(err)
	}
	return  allHooks, status

}

func (db *webhookdb) FindAllRates() ([]Mother, int) {
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		panic(err)

	}
	defer session.Close()
	allFixers := []Mother{}
	err = session.DB(db.dbName).C(db.webhookCollection).Find(nil).All(&allFixers)
	if err != nil{
		status = 0
		log.Println(err)
	}
	return  allFixers, status
}

func(w *webhookdb) FechtAll() int{
	status := 1
	log.Println(Database)
	resp, err := http.Get("http://api.fixer.io/latest")
	if err != nil{
		status = 0
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		status = 0
		panic(err.Error())
	}
	values := Mother{}
	err = json.Unmarshal(body, &values)
	if err != nil{
		status = 0
		panic(err.Error())
	}
	log.Println("FetchAll object",values)
	query, thingy := w.FindFixer(values.Date)
	if values.Date == query.Date{
		log.Println("if 0, then error:  ",thingy)
		log.Println("		todays fixer already exist")
	}else {w.AddFixer(values)}

	resp.Body.Close()
	return status
}

func(db *webhookdb) DropCollection() int {
	status := 1
	session, err := mgo.Dial(db.hostURL)
	if err != nil {
		status = 0
		panic(err)
	}
	defer session.Close()
	session.DB(db.dbName).C(db.webhookCollection).DropCollection()
	return  status
}