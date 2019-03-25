package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"os"
	"time"
)

const(
	giftTypeCoin=iota
	giftTypeCoupon
	giftTypeCouponFix
	giftTypeRealSmall
	giftTypeRealLarge

)

type lotterController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app:=iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})

	initLog()

	return app
}

type gift struct {
	id int
	name string
	pic string
	link string
	gtype int
	data string
	datalist []string
	total int
	left int
	inuse bool
	rate int
	rateMin int
	rateMax int
}

//最大中将号码
const rateMax=10000

var logger *log.Logger

var giftList []*gift

func initLog()  {
	f,_:=os.Create("/var/log/lottery_demo.log")
	logger =log.New(f,"",log.Ldate|log.Lmicroseconds)
}

func initGift()  {
	giftList=make([]*gift,5)

}

func main()  {
	fmt.Println(time.Now().Unix())
}
