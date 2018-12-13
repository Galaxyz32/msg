package messenger

import "time"

type Message struct{
	Index int64 //消息编号
	Time time.Time //消息时间
	Data string //消息内容
	Username string //用户名
}

func NewMessage(preMsg Message,data string,username string)Message{
	newMsg :=Message{}
	newMsg.Index=preMsg.Index + 1
	newMsg.Time=time.Now()
	newMsg.Data=data
	newMsg.Username=username

	return newMsg
}

func WelcomeMsg() Message{
	preMsg:=Message{}
	preMsg.Index=-1
	return NewMessage(preMsg,"Welcome","系统管理员")
}