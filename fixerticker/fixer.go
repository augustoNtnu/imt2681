package fixerticker

import (
	"imt2681/Handlers"
	"time"
	"log"
)


func Ticker(){
var value = 20000
for i:=0; i <value; i++{
	value := Handlers.FixerColl.FechtAll()
	if value == 0 {
		log.Println("error with fetching all")
	}

	Handlers.Database.InvokeAll(Handlers.FixerColl)

	time.Sleep(23 * time.Hour)
	}
}
