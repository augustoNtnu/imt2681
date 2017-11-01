package main

import (
	//"github.com/gorilla/mux"
	"github.com/dchest/uniuri"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"

	"strings"

	"time"
)

type webhookobj struct {
	KeyId 			string		`json:"keyId"`
	WebhookURL      string 		`json:"webhookURL"`
	BaseCurrency    string 		`json:"baseCurrency"`
	TargetCurrency  string 		`json:"targetCurrency"`
	CurrentRate 	float64 	`json:"currentRate"`
	MinTriggerValue float64     `json:"minTriggerValue"`
	MaxTriggerValue float64     `json:"maxTriggerValue"`
}

//var DbPatterns = webhookdb{"127.0.0.1:27017", "local", "patterns"}
var Patterns []string

func HandlerHook(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
//	var status int
	log.Println(len(parts))

	switch req.Method {
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			//status = 500
		}
		//log.Println(string(body))

		t := webhookobj{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			panic(err)
		}

		hash := []byte(uniuri.New())
		w.Write([]byte(hash))
		h := string(hash)
		Patterns = append(Patterns, h)
		t.KeyId = h
		log.Println("object key id: %v", t.KeyId)
		Database.Add(t)

		log.Println(Patterns)
	case "GET":
		db := webhookobj{}
		keyId := parts[2]
		log.Println("keyId for GET", keyId)

		db = Database.find(keyId)
		log.Println(db.KeyId)
		value, err := json.Marshal(db)

		if err != nil {
			log.Println("error encoding webhook:  %v", err.Error())
		}
		w.Write(value)

	case "DELETE":
		keyId := parts[2]

		success := Database.delete(keyId)
		if success == 0 {
			log.Println("error: delete failed")
		}

	default:
		log.Println("error i switch")
	}


}

func HandlerLatest (w http.ResponseWriter, req *http.Request){
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	t := webhookobj{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	time := time.Now().Local().String()
	parts:= strings.Split(time, " ")
	//parts[0] = "2017-10-31"

	//time.Format("2006-01-02")
	log.Println("tid",parts[0])
	log.Println("Kroner for vi hope",t.TargetCurrency)
	resault:= FixerColl.findRates(parts[0])

	ok := resault[t.TargetCurrency]
	ok2,err := json.Marshal(ok)
	if err !=nil {
		log.Println(err)
	}
	w.Write(ok2)





	//thing,err := json.Marshal(resault)
	//if err != nil{
	//	log.Println(err.Error())

	//w.Write(thing)
}

/*
func main (){
	FechtAll()
	//http.HandleFunc("/hreq.URL.Pathello/{patterns}", handlerRegi)
	//http.HandleFunc("/hello", handlerRegi)
	//DbPatterns.retPatt()

	//r:= mux.NewRouter()
	//r.HandleFunc("/hello/", handlerRegi)
	http.HandleFunc("/hello/latest/", HandlerLatest)
	http.HandleFunc("/hello/", HandlerHook)
	http.ListenAndServe("127.0.0.1:8085",nil)

}
*/