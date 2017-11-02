package Handlers

import (
	"testing"
	//"net/http"
	//"net/http/httptest"
	"time"
	"log"

	//"strings"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
)
type testingStruct struct {
	BaseCurrency string `json:"basecurrency"`
	TargetCurrency string `json:"targetcurrency"`


}
var Datatest = webhookdb{"user2:test2@ds042527.mlab.com:42527/cloudtech", "cloudtesting","webhooks"}
var Fixertest = webhookdb{"user2:test2@ds042527.mlab.com:42527/cloudtech","cloudtesting","fixers"}
var testingObj = webhookobj{"dwasdw2d3asd2","localhost:8085/hello", "EUR","NOK",1.6,1.50,2.55}
//var Fixertest = webhookdb{"127.0.0.1:27017", "cloudtest","fixers"}
//var Datatest = webhookdb{"127.0.0.1:27017", "cloudtest","webhooks"}
func TestWebhookdb_Add(t *testing.T) {


	status := Datatest.Add(testingObj)
	if status == 0 {
		t.Error("adding failed")
	}
}

func TestWebhookdb_Count(t *testing.T) {

	value := Datatest.Count()
	if value == 0 {
		t.Error("counting failed")
	}
}
func TestWebhookdb_AddFixer(t *testing.T) {
	Fixertest.FechtAll()

	//fixerDate := time.Now().Local().String()
	//parts := strings.Split(fixerDate, " ")
	fixerdate:= "2017-11-02"

	resault,status  :=Fixertest.FindFixer(fixerdate)
	if status == 0 {
		t.Error("findingFixer failed")
	}
	resault.Date = "2017-11-03"
	status = Fixertest.AddFixer(resault)
	if status == 0 {
		t.Error("addFixer failed")
	}
}

func TestWebhookdb_Find(t *testing.T) {

	resault, status :=Datatest.Find(testingObj.KeyId)
	if status == 0{
		t.Error( "find failed")
	}
	resault.KeyId = ""
}
func TestWebhookdb_FindAllRates(t *testing.T) {

	mordi, value := Fixertest.FindAllRates()
	if value == 0 {
		t.Error("FindAllRates failed")
	}
	mordi[0].Date = ""
}

func TestWebhookdb_FindRates(t *testing.T) {
	fixerDate := time.Now().Local().String()

	value,status :=Fixertest.FindRates(fixerDate)
	if status == 0{
		t.Error("FindRates failed")
	}
	log.Println(value)
}

func TestWebhookdb_FindAll(t *testing.T) {
	value,status := Datatest.FindAll()
	if status == 0{
		t.Error("FindAll failed")
	}
	value[0].KeyId = ""
}

func TestWebhookdb_Delete(t *testing.T) {
	status := Datatest.Delete(testingObj.KeyId)
	if status== 0{
		t.Error("delete failed")
	}
}
func TestWebhookdb_DropCollection(t *testing.T) {
	status := Fixertest.DropCollection()
	if status == 0 {
		t.Error("dropping failed")
	}

}
func TestInvokeAll(t *testing.T) {
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

if status := rr.Code; status != http.StatusOK{
	t.Errorf("handler returned wrong status code:", status, http.StatusOK)
}

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
		t.Errorf("handler not returning 200 ", status,http.StatusOK)
	}
}
