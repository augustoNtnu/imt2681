package main

import(
	"net/http"
	"imt2681/Handlers"
	"os"
)


func main (){

Handlers.FechtAll()

http.HandleFunc("/latest/", Handlers.HandlerLatest)
http.HandleFunc("/app", Handlers.HandlerHook)
http.HandleFunc("/average/",Handlers.HandlerAverage)
http.HandleFunc("/evaluationtrigger/",Handlers.HandlerInvoke)
http.ListenAndServe((":"+os.Getenv("PORT")),nil)
//http.ListenAndServe("127.0.0.1:8080",nil)
}
