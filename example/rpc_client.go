package main

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"net/rpc"
	"github.com/red_envelope/services"
)

func main() {
	c, err := rpc.Dial("tcp", ":18082")
	if err != nil {
		logrus.Panic(err)
	}
	//in := services.RedEnvelopeSendingDTO{
	//	Amount:       decimal.NewFromFloat(1),
	//	UserId:       "1nqMGsA77MzAjDkl8kU5I6u7iww",
	//	Username:     "测试用户",
	//	EnvelopeType: services.GeneralEnvelopeType,
	//	Quantity:     2,
	//	Blessing:     "",
	//}
	//out := &services.RedEnvelopeActivity{}
	//err = c.Call("EnvelopeRpc.SendOut", in, &out)
	//if err != nil {
	//	logrus.Panic(err)
	//}
	//logrus.Infof("%+v", out)
	sendout(c)
	receive(c)
}

func receive(c *rpc.Client) {
	in := services.RedEnvelopeReceiveDTO{
		EnvelopeNo:   "",
		RecvUserId:   "",
		RecvUsername: "",
		AccountNo:    "",
	}
	out := &services.RedEnvelopeItemDTO{}
	err := c.Call("Envelope.Receive", in, out)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("%+v", out)
}
func sendout(c *rpc.Client) {
	in := services.RedEnvelopeSendingDTO{
		Amount:       decimal.NewFromFloat(1),
		UserId:       "1nqMGsA77MzAjDkl8kU5I6u7iww",
		Username:     "测试用户",
		EnvelopeType: services.GeneralEnvelopeType,
		Quantity:     2,
		Blessing:     "",
	}
	out := &services.RedEnvelopeActivity{}
	err := c.Call("EnvelopeRpc.SendOut", in, &out)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("%+v", out)
}

