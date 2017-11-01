package main

type webhookobj struct {
	KeyId 			string		`json:"keyId"`
	WebhookURL      string 		`json:"webhookURL"`
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
var Database = webhookdb{"127.0.0.1:27017","local","webhooks"}
var FixerColl = webhookdb{"127.0.0.1:27017","local","fixerrates"}

type webhookdb struct {
	hostURL           string
	dbName            string
	webhookCollection string
}