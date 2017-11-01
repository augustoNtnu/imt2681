package main

import(
	"net/http"
	"imt2681/Handlers"
	"os"
)


func main (){

Handlers.FechtAll()

http.HandleFunc("/latest/", Handlers.HandlerLatest)
http.HandleFunc("/", Handlers.HandlerHook)
http.HandleFunc("/average/",Handlers.HandlerAverage)
http.HandleFunc("/evaluationtrigger/",Handlers.HandlerInvoke)
http.ListenAndServe((":"+os.Getenv("PORT")),nil)
}
