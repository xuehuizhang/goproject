package util

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type TailObj struct {
	Tail *tail.Tail
	Conf CollectConf
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailObjMgr struct {
	Tails []*TailObj
	MsgChan chan *TextMsg
}
var (
	tailObjMgr *TailObjMgr
)
func InitTail(conf []CollectConf,chanSize int)(err error){
	fmt.Println(chanSize)
	if len(conf)==0{
		fmt.Println("conf err")
		err=fmt.Errorf("invalid config for log collect,conf:%v",conf)
		return
	}
	tailObjMgr=&TailObjMgr{
		MsgChan:make(chan *TextMsg,chanSize),
	}
	tailObjMgr.Tails=make([]*TailObj,1)
	for _,v:=range conf {
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})
		if err != nil {
			fmt.Println("tail file err", err)
			err=errTail
			return
		}
		tailObj:=TailObj{}
		tailObj.Conf=v
		tailObj.Tail=tails
		tailObjMgr.Tails=append(tailObjMgr.Tails,&tailObj)

		go readFromTail(&tailObj)
	}
	return
}
func readFromTail(obj *TailObj)  {
	for true{
		msg,ok:=<-obj.Tail.Lines
		if !ok{
			logs.Error("tail lines err",obj.Tail.Filename)
			time.Sleep(time.Millisecond*100)
			continue
		}
		TextMsg:=TextMsg{
			Msg:msg.Text,
			Topic:obj.Conf.Topic,
		}
		tailObjMgr.MsgChan<-&TextMsg
	}
}

func GetOneLine()(msg *TextMsg){
	msg=<-tailObjMgr.MsgChan
	return
}