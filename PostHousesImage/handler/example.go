package handler

import (
	"context"
	example "IHome/PostHousesImage/proto/example"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
)

type Example struct{}

func (e *Example) PostHousesImage(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("发送房屋图片PostHousesImage  /api/v1.0/houses/:id/images")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	return nil
}
