package main

import "net/http"


func main (){
FechtAll()
//http.HandleFunc("/hreq.URL.Pathello/{patterns}", handlerRegi)
//http.HandleFunc("/hello", handlerRegi)
//DbPatterns.retPatt()

//r:= mux.NewRouter()
//r.HandleFunc("/hello/", handlerRegi)
http.HandleFunc("/hello/latest/", HandlerLatest)
http.HandleFunc("/hello/", HandlerHook)
http.ListenAndServe("127.0.0.1:8085",nil)

}
