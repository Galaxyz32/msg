package main

import "messenger"

func main(){
	msg:=messenger.NewMsgList()
//	msg.Print()
	println()

	msg.SendMsg("早上好","小明")

	msg.SendMsg("你也好","小红")
	msg.Print()


}



