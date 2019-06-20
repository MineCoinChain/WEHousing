package handler

import (
	"context"

	example "IHome/DeleteSession/proto/example"
	"IHome/IhomeWeb/utils"
	"fmt"
	"log"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) DeleteSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" DELETE session    /api/v1.0/session !!!")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis connect error")
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(rsp.Errno)
		return nil
	}
	sessionidname :=  req.Sessionid + "name"
	sessioniduserid :=  req.Sessionid + "user_id"
	sessionidmobile :=  req.Sessionid + "mobile"

	//从缓存中获取session 那么使用唯一识别码
	bm.Delete(sessionidname)
	bm.Delete(sessioniduserid)
	bm.Delete(sessionidmobile)

	return nil
}
