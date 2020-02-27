package controller

import (
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/JabinGP/ssr-sub-hub/model"

	"github.com/kataras/iris/v12"
)

// GetUserConfig 获取配置
func GetUserConfig(ctx iris.Context) {
	var req struct {
		Username string
		Password string
	}
	var res struct {
		ConfigString string
	}

	req.Username = ctx.Params().Get("username")
	req.Password = ctx.Params().Get("password")

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

	exist := true
	userConfigName := "./user/" + req.Username + ".toml"
	if _, err := os.Stat(userConfigName); os.IsNotExist(err) {
		exist = false
	}
	if !exist {
		res.ConfigString = "# 初始化于" + time.Now().Format("2006-01-02 15:04:05") + "，请编辑后保存。"
		ctx.JSON(res)
		return
	}

	userConfigByteList, err := ioutil.ReadFile(userConfigName)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorReadUserConfig(err))
		return
	}
	res.ConfigString = string(userConfigByteList)
	ctx.JSON(res)
}

// PostUserConfig ...
func PostUserConfig(ctx iris.Context) {
	var req struct {
		Username string
		Password string
		Config   string
	}
	ctx.ReadJSON(&req)

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

	if len(req.Config) < 5 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorBadUserConfig(errors.New("上传的配置文件过短")))
		return
	}

	err = userService.UpdateUserConfig(req.Username, req.Config)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorUpdateUserConfig(err))
		return
	}

	ctx.JSON("")
}
