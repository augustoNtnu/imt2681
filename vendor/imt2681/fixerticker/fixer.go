package fixerticker

import (
	"imt2681/Handlers"
	"net/http"

	"log"
//	"encoding/json"

	"io/ioutil"
	"encoding/json"
)


func ticker(){
var value = 20000
for i:=0; i <value; i++{
	Handlers.FechtAll()
	Handlers.InvokeAll

	time.Sleep(23 * time.Hour)
	}
}