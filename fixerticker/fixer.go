package fixerticker

import (
	"imt2681/Handlers"
	"net/http"
	"log"
//	"encoding/json"

	"io/ioutil"
	"encoding/json"
)

func FechtAll(){
	log.Println(Database)
	resp, err := http.Get("http://api.fixer.io/latest")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err.Error())
	}
	values := Mother{}
	err = json.Unmarshal(body, &values)
	if err != nil{
		panic(err.Error())
	}
	log.Println(values)
	Handlers.FixerColl.
}