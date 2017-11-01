package Handlers

import (
	"github.com/dchest/uniuri"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
	"time"
	"bytes"


)





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
		t.KeyId = h
		log.Println("object key id: %v", t.KeyId)
		Database.Add(t)


	case "GET":
		db := webhookobj{}
		keyId := parts[2]
		log.Println("keyId for GET", keyId)

		db = Database.Find(keyId)
		log.Println(db.KeyId)
		value, err := json.Marshal(db)

		if err != nil {
			log.Println("error encoding webhook:  %v", err.Error())
		}
		w.Write(value)

	case "DELETE":
		keyId := parts[2]

		success := Database.Delete(keyId)
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
	time := time.Now().UTC().String()
	parts:= strings.Split(time, " ")
	parts[0] = "2017-10-31"

	//time.Format("2006-01-02")
	log.Println("tid",parts[0])
	log.Println("Kroner for vi hope",t.TargetCurrency)
	resault:= FixerColl.FindRates(parts[0])

	ok := resault[t.TargetCurrency]
	ok2,err := json.Marshal(ok)
	if err !=nil {
		log.Println(err)
	}
	w.Write(ok2)
}

func HandlerInvoke(w http.ResponseWriter, req *http.Request) {
	webhooks := Database.FindAll()


	timeValue := time.Now().Local().String()
	parts := strings.Split(timeValue, " ")
	rates := FixerColl.FindRates(parts[0])
	nrOfWebhooks := len(webhooks) - 1
	path := strings.Split(req.URL.Path, "/")

	for i := 0; i <= nrOfWebhooks; i++ {
		webhooks[i].CurrentRate = rates[webhooks[i].TargetCurrency]
		if  path[2] == "evaluationtrigger" {
			body, err := json.Marshal(webhooks[i])

			if err != nil {
				log.Println(err)
				} else {

				response, err := http.Post(webhooks[i].WebhookURL, "application/json", bytes.NewBuffer(body))
				defer response.Body.Close()
				if err != nil {
					log.Println(err)
				}
				if response.StatusCode != 200 || response.StatusCode != 204 {
					log.Println("Invoking failed")
				}

			}
		}
	}
}

func HandlerAverage(w http.ResponseWriter, req *http.Request){
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var average float64
	baseValues := webhookobj{}
	err = json.Unmarshal(body,baseValues)
	allRates := FixerColl.FindAllRates()
	//length := len(allRates)-1

	for i := 0; i <= 2; i++{

		average += allRates[i].Rates[baseValues.TargetCurrency]
	}
	response, err := json.Marshal(average)
	if err != nil{
		log.Println(err)
	}

	w.Write(response)
}


func InvokeAll(){
	webhooks := Database.FindAll()


	timeValue := time.Now().Local().String()
	parts := strings.Split(timeValue, " ")
	rates := FixerColl.FindRates(parts[0])
	nrOfWebhooks := len(webhooks) - 1


	for i := 0; i <= nrOfWebhooks; i++ {
		currentWebRate := rates[webhooks[i].TargetCurrency]
		webhooks[i].CurrentRate = rates[webhooks[i].TargetCurrency]
		if webhooks[i].MaxTriggerValue > currentWebRate || webhooks[i].MinTriggerValue < currentWebRate {
			body, err := json.Marshal(webhooks[i])
			if err != nil {
				log.Println(err)
			} else {

				response, err := http.Post(webhooks[i].WebhookURL, "application/json", bytes.NewBuffer(body))
				defer response.Body.Close()
				if err != nil {
					log.Println(err)
				}
				if response.StatusCode != 200 || response.StatusCode != 204 {
					log.Println("Invoking failed")
				}

			}
		}
	}
}
