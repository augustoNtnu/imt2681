package main

import(
	"net/http"
	"imt2681/Handlers"
)


func main (){
Handlers.FechtAll()

http.HandleFunc("/hello/latest/", Handlers.HandlerLatest)
http.HandleFunc("/hello/", Handlers.HandlerHook)
http.HandleFunc("/hello/average/",Handlers.HandlerAverage)
http.HandleFunc("/hello/evaluationtrigger/",Handlers.HandlerInvoke)
http.ListenAndServe("127.0.0.1:8085",nil)
}
