package web

import (
	"github.com/ABottomCoder/infra"
	"github.com/ABottomCoder/infra/base"
	"github.com/kataras/iris/v12"
	"github.com/red_envelope/services"
)

func init() {
	infra.RegisterApi(&EnvelopeApi{})
}

type EnvelopeApi struct {
	service services.RedEnvelopeService
}

func (e *EnvelopeApi) Init() {
	e.service = services.GetRedEnvelopeService()
	groupRouter := base.Iris().Party("/v1/envelope")
	groupRouter.Post("/sendout", e.sendOutHandler)
	groupRouter.Post("/receive", e.receiveHandler)

	groupRouter.Get("/getReceived", e.getReceivedHandler)
	groupRouter.Get("/getReceivable", e.getReceivableHandler)

}

func (e *EnvelopeApi) receiveHandler(ctx iris.Context) {
	dto := services.RedEnvelopeReceiveDTO{}
	err := ctx.ReadJSON(&dto)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	item, err := e.service.Receive(dto)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	r.Data = item
	ctx.JSON(r)
}

//{
//	"envelopeType": 0,
//	"username": "",
//	"userId": "",
//	"blessing": "",
//	"amount": "0",
//	"quantity": 0
//}
func (e *EnvelopeApi) sendOutHandler(ctx iris.Context) {
	dto := services.RedEnvelopeSendingDTO{}
	err := ctx.ReadJSON(&dto)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	activity, err := e.service.SendOut(dto)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	r.Data = activity
	ctx.JSON(r)

}

func (e *EnvelopeApi) getReceivedHandler(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	page, err1 := ctx.URLParamInt("page")
	size, err2 := ctx.URLParamInt("size")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" || err1 != nil || err2 != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "参数错误"
		ctx.JSON(&r)
		return
	}
	items := e.service.ListReceived(userId, page, size)
	r.Data = items
	ctx.JSON(r)
}

func (e *EnvelopeApi) getReceivableHandler(ctx iris.Context) {
	page, err1 := ctx.URLParamInt("page")
	size, err2 := ctx.URLParamInt("size")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err1 != nil || err2 != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "参数错误"
		ctx.JSON(&r)
		return
	}
	orders := e.service.ListReceivable(page, size)
	r.Data = orders
	ctx.JSON(r)
}
