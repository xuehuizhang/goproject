package model

import "github.com/hpcloud/tail"

type TailObj struct {
	Tail *tail.Tail
	Conf CollectConf
}

type TailObjMgr struct {
	Tails []*TailObj
	MsgChan chan *TextMsg
}

type TextMsg struct {
	Msg string
	Topic string
}