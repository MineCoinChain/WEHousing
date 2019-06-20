package handler

import (
	"context"
	example "IHome/GetSession/proto/example"
	"IHome/IhomeWeb/utils"
	"log"
	"github.com/garyburd/redigo/redis"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) GetSession(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	if err != nil {
		log.Println("redis连接失败")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	sessionid := req.Sessionid
	namekey := sessionid + "name"
	value := bm.Get(namekey)
	name, _ := redis.String(value, nil)
	rsp.Data = name
	return nil
}
