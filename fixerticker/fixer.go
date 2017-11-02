package fixerticker

import (
	"imt2681/Handlers"
	"time"
)


func Ticker(){
var value = 20000
for i:=0; i <value; i++{
	Handlers.FechtAll()
	Handlers.InvokeAll()

	time.Sleep(23 * time.Hour)
	}
}
