package messenger

import "fmt"

type MsgList struct {
	Messages []*Message
}

func NewMsgList() *MsgList{
	Welmsg:=WelcomeMsg()
	messagelist:=MsgList{}
	messagelist.AppendMsg(&Welmsg)
	return &messagelist
}

func (msg *MsgList) AppendMsg(newMsg *Message){
	msg.Messages=append(msg.Messages,newMsg)
}

func (msg *MsgList) SendMsg(data string,username string) {
	preMsg := msg.Messages[len(msg.Messages)-1]
	newMsg := NewMessage(*preMsg,data,username)
	msg.AppendMsg(&newMsg)
}

func (msg *MsgList) Print(){
	for _,msg:=range msg.Messages{
		fmt.Printf("%s：\n",msg.Username)
		fmt.Printf("编号：%d\n",msg.Index)
		fmt.Println("时间："+msg.Time.Format("02 Jan 2006 15:04"))
		fmt.Printf("内容：%s\n",msg.Data)
		fmt.Println()
	}
}