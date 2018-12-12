package beego_mailer

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func InitMailerQueue(ctx *context.Context) {
}

func ProcessMailerQueue(ctx *context.Context) {
}

func InitMailerFilter() {

	beego.InsertFilter("*", beego.BeforeRouter, InitMailerQueue)
	beego.InsertFilter("*", beego.FinishRouter, ProcessMailerQueue, false)

	beego.Info("Keenio filter initialized")

}
