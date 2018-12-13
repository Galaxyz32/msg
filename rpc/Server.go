package main

import (
	"encoding/json"
	"io"
	"messenger"
	"net/http"
)

var msglist *messenger.MsgList

func run(){
	http.HandleFunc("/msg/recv",MsgGetHandler)
	http.HandleFunc("/msg/send",MsgSendHandler)
	http.ListenAndServe("localhost:8888",nil)
}

func MsgGetHandler(w http.ResponseWriter,r *http.Request){
	bytes,error:=json.Marshal(msglist)
	if error!=nil{
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

func MsgSendHandler(w http.ResponseWriter,r *http.Request){
	msgdata:=r.URL.Query().Get("msg")
	username:=r.URL.Query().Get("id")
	msglist.SendMsg(msgdata,username)
	MsgGetHandler(w,r)
}

func main(){
	msglist =messenger.NewMsgList()
	run()
}