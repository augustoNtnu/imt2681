package Handlers

import (
	"testing"
	//"net/http"
	//"net/http/httptest"
	//"time"
	"log"

	//"strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	//"strings"
	"strings"
	"time"

)
type testingStruct struct {
	BaseCurrency string `json:"basecurrency"`
	TargetCurrency string `json:"targetcurrency"`


}
//var Datatest = webhookdb{"user2:test2@ds042417.mlab.com:42417/cloudtesting", "cloudtesting","webhooks"}
//var Fixertest = webhookdb{"user2:test2@ds042417.mlab.com:42417/cloudtesting","cloudtesting","fixers"}
var testingObj = webhookobj{"dwasdw2d3asd2","http://www.google.com/", "EUR","NOK",1.46,1.50,2.55}
var Fixertest = webhookdb{"mongodb://localhost", "cloudtest","fixers"}
var Datatest = webhookdb{"mongodb://localhost", "cloudtest","webhooks"}

func TestWebhookdb_Add(t *testing.T) {

	status := Datatest.Add(testingObj)
	if status == 0 {
		t.Error("adding failed")
	}
	log.Println("Add finished")
}

func TestWebhookdb_Count(t *testing.T) {

	value := Datatest.Count()
	if value == 0 {
		t.Error("counting failed")
	}
	log.Println("Count finshed")
}
func TestWebhookdb_AddFixer(t *testing.T) {
	Fixertest.FechtAll()

	fixerDate := time.Now().Local().String()
	parts := strings.Split(fixerDate, " ")

	resault,status  :=Fixertest.FindFixer(parts[0])
	if status == 0 {
		t.Error("findingFixer failed")
	}
	resault.Date = "2017-11-03"
	status = Fixertest.AddFixer(resault)
	if status == 0 {
		t.Error("addFixer failed")
	}
	log.Println("addfixer and find fixer done")
}

func TestWebhookdb_Find(t *testing.T) {

	resault, status :=Datatest.Find(testingObj.KeyId)
	if status == 0{
		t.Error( "find failed")
	}
	resault.KeyId = ""
	log.Println("Find finished")
}
func TestWebhookdb_FindAllRates(t *testing.T) {

	mordi, value := Fixertest.FindAllRates()
	if value == 0 {
		t.Error("FindAllRates failed")
	}
	mordi[0].Date = ""
log.Println("findAllRates finished")
}

func TestWebhookdb_FindRates(t *testing.T) {
	temp := time.Now().Local().String()
	fixerDate := strings.Split(temp, " ")
	log.Println("fixerDate: ", fixerDate)

	value,status :=Fixertest.FindRates(fixerDate[0])
	if status == 0{
		t.Error("FindRates failed")
	}
	log.Println(value,"Find rates finished")
}

func TestWebhookdb_FindAll(t *testing.T) {
	value,status := Datatest.FindAll()
	if status == 0{
		t.Error("FindAll failed")
	}
	value[0].KeyId = ""
	log.Println("FindAll finished")
}

func TestWebhookdb_InvokeAll(t *testing.T) {
	Datatest.InvokeAll(Fixertest)
}

func TestHandlerInvoke(t *testing.T) {
	req, err := http.NewRequest("POST", "/evaluationtrigger/",nil)
	if err != nil{
	t.Fatal(err)
	}
rr:= httptest.NewRecorder()
handler := http.HandlerFunc(HandlerInvoke)
handler.ServeHTTP(rr,req)
log.Println("testHandlerInvoke Finished")
}

func TestHandlerAverage(t *testing.T) {
	test := testingStruct{"EUR", "NOK"}
	body, err:= json.Marshal(test)
	if err != nil{
		t.Error("marshalling failed")
	}

	req, err := http.NewRequest("POST", "/evaluationtrigger/",bytes.NewBuffer(body))
	if err != nil {
		t.Error("nw request faield")
	}

	rr:=httptest.NewRecorder()
	handler := http.HandlerFunc(HandlerAverage)
	handler.ServeHTTP(rr,req)
	if status:= rr.Code; status != http.StatusOK{
		t.Error("handler not returning 200 ", status,http.StatusOK)
	}
}
func TestWebhookdb_DropCollection(t *testing.T) {
	status := Fixertest.DropCollection()
	if status == 0 {
		t.Error("dropping failed")
	}


}
func TestWebhookdb_Delete(t *testing.T) {
	status := Datatest.Delete(testingObj.KeyId)
	if status== 0{
		t.Error("delete failed")
	}
}