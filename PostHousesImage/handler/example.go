package handler

import (
	"context"
	example "IHome/PostHousesImage/proto/example"
	"github.com/astaxie/beego"
	"IHome/IhomeWeb/utils"
	"path"
	"IHome/IhomeWeb/model"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Example struct{}

func (e *Example) PostHousesImage(ctx context.Context, req *example.Request, rsp *example.Response) error {
	beego.Info("发送房屋图片PostHousesImage  /api/v1.0/houses/:id/images")
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	/*获取文件的后缀名*/
	fileExt := path.Ext(req.Filename)

	/*将获取到的图片数据成为二进制信息存入fastdfs*/
	fileId, err := utils.Uploadbybuf(req.Image, fileExt[1:])
	if err != nil {
		beego.Info("Postupavatar  models.UploadByBuffer err", err)
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	/*获取house_id*/
	houseId, _ := strconv.Atoi(req.Id)
	//创建house 对象
	House := models.House{Id: houseId}
	//创建数据库句柄
	o := orm.NewOrm()
	err = o.Read(&House)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	/*判断index_image_url 是否为空    一般第一张图片设置为主要图片*/
	if House.Index_image_url == "" {
		House.Index_image_url = fileId
	}
	/*将该图片添加到 house的全部图片当中*/
	houseImage := models.HouseImage{House: &House, Url: fileId}
	House.Images = append(House.Images, &houseImage)
	//将图片对象插入表单之中
	_, err = o.Insert(&houseImage)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	//对house表进行更新
	_, err = o.Update(&House)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	rsp.Url = fileId
	return nil
}
