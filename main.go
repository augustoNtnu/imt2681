package main

import "net/http"


func main (){
FechtAll()

http.HandleFunc("/hello/latest/", HandlerLatest)
http.HandleFunc("/hello/", HandlerHook)
http.HandleFunc("/hello/average/",handlerAverage)
http.HandleFunc("/hello/evaluationtrigger/",handlerInvoke)
http.ListenAndServe("127.0.0.1:8085",nil)
}
