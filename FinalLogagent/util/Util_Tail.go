package util

import (
	"FinalLogagent/model"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)
var (
	tailObjMgr model.TailObjMgr
)
func InitTail(confs []model.CollectConf,chanSize int)(err error){
	if len(confs)==0{
		err=fmt.Errorf("invalid config for log collect,%v",confs)
		return
	}

	tailObjMgr=model.TailObjMgr{
		MsgChan:make(chan *model.TextMsg,chanSize),
	}
	for _,v:=range confs{
		tail,errTail:=tail.TailFile(v.LogPath,tail.Config{
			ReOpen:true,
			Follow:true,
			MustExist:false,
			Poll:true,
		})
		if err!=nil{
			err=errTail
			return
		}
		tailObj:=&model.TailObj{Conf:v}
		tailObj.Tail=tail

		tailObjMgr.Tails=append(tailObjMgr.Tails,tailObj)

		go readFromTail(tailObj)
	}
	return
}

func readFromTail(tail *model.TailObj){
	for{
		msg,ok:=<-tail.Tail.Lines
		if !ok{
			logs.Error("tail lines err",tail.Tail.Filename)
			time.Sleep(time.Millisecond*100)
			continue
		}
		textMsg:=model.TextMsg{
			msg.Text,
			tail.Conf.Topic,
		}
		tailObjMgr.MsgChan<-&textMsg
	}
}

func GetOneLine()(msg *model.TextMsg){
	msg=<-tailObjMgr.MsgChan
	return
}