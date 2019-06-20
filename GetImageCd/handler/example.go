package handler

import (
	"context"

	example "IHome/GetImageCd/proto/example"
	"github.com/afocus/captcha"
	"image/color"
	"IHome/IhomeWeb/utils"
	"time"
)

type Example struct{}

func (e *Example) GetImageCd(ctx context.Context, req *example.Request, rsp *example.Response) error {
	///********1.初始化返回值*******/
	//rsp.Errno = utils.RECODE_OK
	//rsp.Errmsg = utils.RecodeText(rsp.Errno)
	///********2.生成随机数和图片*******/
	////创建1个句柄
	//cap := captcha.New()
	////通过句柄调用 字体文件
	//if err := cap.SetFont("comic.ttf"); err != nil {
	//	panic(err.Error())
	//}
	////设置图片的大小
	//cap.SetSize(91, 41)
	//// 设置干扰强度
	//cap.SetDisturbance(captcha.MEDIUM)
	//// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	////SetFrontColor(colors ...color.Color)  这两个颜色设置的函数属于不定参函数
	//cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//// 设置背景色 可以多个 随机替换背景色 默认白色
	//cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	////生成图片 返回图片和 字符串(图片内容的文本形式)
	//img, str := cap.Create(4, captcha.NUM)
	///********3.获取uuid*******/
	//uuid := req.Uuid
	///********4.连接redis*******/
	//bm, err := utils.RedisOpen(utils.G_server_name, utils.G_redis_addr, utils.G_redis_port, utils.G_redis_dbnum)
	//if err != nil {
	//	log.Println("redis连接错误", err)
	//	rsp.Errno = utils.RECODE_DBERR
	//	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	//	return nil
	//}
	///********5.存入数据*******/
	//bm.Put(uuid, str, time.Second*300)
	///********6.将图片拆分，赋值到proto*******/
	////6 将图片拆分 赋值到proto
	//a := *img
	//
	//b := *(a.RGBA)
	//
	////pix
	//for _, value := range b.Pix {
	//
	//	rsp.Pix  =append(rsp.Pix ,uint32(value))
	//}
	//
	//
	////stride
	//rsp.Stride = int64(b.Stride)
	//
	////point
	//rsp.Min = &example.ResponsePoint{X:int64(b.Rect.Min.X),
	//	Y:int64(b.Rect.Min.Y)}
	//rsp.Max = &example.ResponsePoint{X:int64(b.Rect.Max.X),
	//	Y:int64(b.Rect.Max.Y)}
	//
	//
	//return nil
	//1 初始化返回值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)


	//2 生成随即数与图片
	//创建一个句柄

	cap := captcha.New()
	//设置字体
	if err := cap.SetFont("comic.ttf"); err != nil {
		//抛出异常
		panic(err.Error())
	}

	//设置突破大小
	cap.SetSize(90, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255},color.RGBA{0, 0, 0, 200})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})


	//随即生存 图片 与 验证码
	img, str := cap.Create(4, captcha.NUM)

	//打印字符串
	println(str)

	//3 获取uuid
	uuid := req.Uuid

	//4 连接redis
	bm ,err :=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,
		utils.G_redis_port,utils.G_redis_dbnum)
	if err !=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}

	//5 存入数据
	bm.Put(uuid ,str ,time.Second * 300)


	//6 将图片拆分 赋值到proto
	a := *img

	b := *(a.RGBA)

	//pix
	for _, value := range b.Pix {

		rsp.Pix  =append(rsp.Pix ,uint32(value))
	}


	//stride
	rsp.Stride = int64(b.Stride)

	//point
	rsp.Min = &example.ResponsePoint{X:int64(b.Rect.Min.X),
		Y:int64(b.Rect.Min.Y)}
	rsp.Max = &example.ResponsePoint{X:int64(b.Rect.Max.X),
		Y:int64(b.Rect.Max.Y)}

	return nil
}
