package controller

import (
	"errors"
	"log"

	"github.com/JabinGP/ssr-sub-hub/model"
	"github.com/JabinGP/ssr-sub-hub/model/reqo"

	"github.com/kataras/iris/v12"
)

// GetSSR 获取SSR订阅文件
func GetSSR(ctx iris.Context) {
	req := reqo.GetSub{}
	ctx.ReadQuery(&req)

	if req.Username == "" || req.Password == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户名密码不能为空")))
		return
	}

	err := userService.Verify(req.Username, req.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
		return
	}

	linkList := userService.GetSSRLinkList(req.Username)
	log.Println(linkList)
	data, err := subService.GetAllSSR(linkList)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorGenerateFile(err))
		return
	}

	ctx.ContentType("application/octet-stream; charset=utf-8")
	ctx.Header("Content-Disposition", "attachment; filename=sub.txt")
	ctx.Write(data)
}

// GetSSRList 获取SSR订阅链接列表
func GetSSRList(ctx iris.Context) {
	req := reqo.GetSub{}
	ctx.ReadQuery(&req)

	if req.Username == "" || req.Password == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户名密码不能为空")))
		return
	}

	err := userService.Verify(req.Username, req.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(err))
		return
	}

	linkList := userService.GetSSRLinkList(req.Username)

	ctx.JSON(linkList)
}
