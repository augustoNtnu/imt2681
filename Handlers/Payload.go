package Handlers

import "os"

//import "os"

type webhookobj struct {
	KeyId 			string		`json:"keyId,omitempty"`
	WebhookURL      string 		`json:"webhookURL,omitempty"`
	BaseCurrency    string 		`json:"baseCurrency"`
	TargetCurrency  string 		`json:"targetCurrency"`
	CurrentRate 	float64 	`json:"currentRate"`
	MinTriggerValue float64     `json:"minTriggerValue"`
	MaxTriggerValue float64     `json:"maxTriggerValue"`
}
type invokeHook struct{

	BaseCurrency    string 		`json:"baseCurrency"`
	TargetCurrency  string 		`json:"targetCurrency"`
	CurrentRate 	float64 	`json:"currentRate"`
	MinTriggerValue float64     `json:"minTriggerValue"`
	MaxTriggerValue float64     `json:"maxTriggerValue"`
}

type Mother struct {
	Base string `json:"base"`
	Date string	`json:"date"`

	Rates map[string]float64 `json:"rates"`
}
//mongodb://<dbuser>:<dbpassword>@ds229415.mlab.com:29415/assigment2
//var Database = webhookdb{("mongodb://"+os.Getenv("userName")+":"+os.Getenv("userPass")+"@ds229415.mlab.com:29415/assigment2"),os.Getenv("Database"),os.Getenv("COLLECTION1")}
var Database = webhookdb{("mongodb://user:test@ds229415.mlab.com:29415/assigment2"),"assigment2","webhooks"}

var FixerColl = webhookdb{("mongodb://"+os.Getenv("userName")+":"+os.Getenv("userPass")+"@ds229415.mlab.com:29415/assigment2"),os.Getenv("Database"),os.Getenv("COLLECTION2")}

type webhookdb struct {
	hostURL           string
	dbName            string
	webhookCollection string

}


//http.Error(w, http.StatusText(status), status)