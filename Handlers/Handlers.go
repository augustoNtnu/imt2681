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
	status := 200
//	var status int
	log.Println(len(parts))

	switch req.Method {
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			status = 400
		}
		//log.Println(string(body))

		t := webhookobj{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			status = 400
		}

		hash := []byte(uniuri.New())
		w.Write([]byte(hash))
		h := string(hash)
		t.KeyId = h
		log.Println("object key id: %v", t.KeyId)
		value2 := Database.Add(t)
		if value2 == 0{
			status = 500
		}


	case "GET":
		db := webhookobj{}
		keyId := parts[2]
		var faenskap int
		log.Println("keyId for GET", keyId)

		db, faenskap = Database.Find(keyId)
		if faenskap == 0 {
			status = 500
		}
		log.Println(db.KeyId)
		value, err := json.Marshal(db)
		if err != nil {
			log.Println("error encoding webhook:  %v", err.Error())
			status = 500
		}
		w.Write(value)

	case "DELETE":
		keyId := parts[2]

		success := Database.Delete(keyId)
		if success == 0 {
			log.Println("error: delete failed")
			status = 404
		}

	default:
		log.Println("error i switch")
	}
	//http.Error(w, http.StatusText(status), status)
	w.WriteHeader(status)

}

func HandlerLatest (w http.ResponseWriter, req *http.Request){
	status := 200
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
		status = 400
	}
	t := webhookobj{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
		status = 400
	}
	time := time.Now().UTC().String()
	parts:= strings.Split(time, " ")
	parts[0] = "2017-11-01"

	//time.Format("2006-01-02")
	log.Println("tid",parts[0])
	log.Println("Kroner for vi hope",t.TargetCurrency)
	resault, thingy:= FixerColl.FindRates(parts[0])
	if thingy == 0{
		status = 500
	}

	ok := resault[t.TargetCurrency]
	ok2,err := json.Marshal(ok)
	if err !=nil {
		log.Println(err)
		status = 500
	}
	w.Write(ok2)
	w.WriteHeader(status)
}

func HandlerInvoke(w http.ResponseWriter, req *http.Request) {
	status := 200
	webhooks, thingy := Database.FindAll()
	if thingy == 0{
		status = 500
	}

	timeValue := time.Now().Local().String()
	parts := strings.Split(timeValue, " ")
	rates, thang := FixerColl.FindRates(parts[0])
	if thang == 0 {
		status = 500
	}
	nrOfWebhooks := len(webhooks) -1
	path := strings.Split(req.URL.Path, "/")

	for i := 0; i <= nrOfWebhooks; i++ {
		webhooks[i].CurrentRate = rates[webhooks[i].TargetCurrency]
		webhooks[i].WebhookURL = ""
		webhooks[i].KeyId= ""
		if  path[2] == "evaluationtrigger" {
			body, err := json.Marshal(webhooks[i])

			if err != nil {
				log.Println(err)
				status = 500
				} else {

				response, err := http.Post(webhooks[i].WebhookURL, "application/json", bytes.NewBuffer(body))
				defer response.Body.Close()
				if err != nil {
					log.Println(err)
					status = 500
				}
				if response.StatusCode != 200 || response.StatusCode != 204 {
					log.Println("Invoking failed")
				}

			}
		}
	}
	w.WriteHeader(status)
}

func HandlerAverage(w http.ResponseWriter, req *http.Request){
	status := 200
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
		status = 400
	}
	var average float64
	baseValues := webhookobj{}
	err = json.Unmarshal(body,baseValues)
	if err != nil{
		status = 500
	}
	allRates, thingy := FixerColl.FindAllRates()
	if thingy == 0 {
		status = 500
	}
	length := len(allRates)

	for i,rate := range allRates{

		average +=rate.Rates[baseValues.TargetCurrency]
		if i == 2 {break}
	}

	average /=float64(length)
	response, err := json.Marshal(average)
	if err != nil{
		log.Println(err)
		status = 500
	}

	w.Write(response)
	w.WriteHeader(status)
}


func (db *webhookdb)InvokeAll(fixer webhookdb) {
	webhooks, thingy := db.FindAll()
	if thingy == 0 {
		log.Println("findAll failed")
	} else {

		timeValue := time.Now().Local().String()
		parts := strings.Split(timeValue, " ")
		//parts := "2017-11-10"
		rates, thang := fixer.FindRates(parts[0])
		if thang == 0 {
			log.Println("findRates failed")
		}else {
			nrOfWebhooks := 0
			log.Println("# of webhooks",nrOfWebhooks)

			for i := 0; i <= nrOfWebhooks; i++ {
				currentWebRate := rates[webhooks[i].TargetCurrency]
				webhooks[i].CurrentRate = rates[webhooks[i].TargetCurrency]
				tempUrl := webhooks[i].WebhookURL
				webhooks[i].WebhookURL = ""
				webhooks[i].KeyId= ""
				if webhooks[i].MaxTriggerValue > currentWebRate || webhooks[i].MinTriggerValue < currentWebRate {
					body, err := json.Marshal(webhooks[i])
					if err != nil {
						log.Println(err)
					} else {

						response, err := http.Post(tempUrl, "application/json", bytes.NewBuffer(body))

						if err != nil {
							log.Println(err)
						}
						if response.StatusCode != 200 || response.StatusCode != 204 {
							log.Println("Invoking failed")
						}
						response.Body.Close()
					}
				}
			}
		}

	}
}
