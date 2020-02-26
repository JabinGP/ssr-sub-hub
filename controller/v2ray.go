package controller

import (
	"errors"
	"log"

	"github.com/JabinGP/ssr-sub-hub/model"

	"github.com/kataras/iris/v12"
)

// GetV2ray Get "/v2ray"
func GetV2ray(ctx iris.Context) {
	var req struct {
		Username string
		Password string
	}
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

	v2rayList, err := userService.GetV2rayList(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorReadUserConfig(err))
		return
	}
	log.Println(v2rayList)

	var linkList []string
	for _, v2ray := range v2rayList {
		if v2ray.Enable {
			linkList = append(linkList, v2ray.URL)
		}
	}
	log.Println(linkList)

	data, err := subService.GetAllV2ray(linkList)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorGenerateFile(err))
		return
	}

	ctx.ContentType("application/octet-stream; charset=utf-8")
	ctx.Header("Content-Disposition", "attachment; filename=sub.txt")
	ctx.Write(data)
}

// GetV2rayList Get "/v2ray/list"
func GetV2rayList(ctx iris.Context) {
	var req struct {
		Username string
		Password string
	}
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

	v2rayList, err := userService.GetV2rayList(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorReadUserConfig(err))
		return
	}

	ctx.JSON(v2rayList)
}
