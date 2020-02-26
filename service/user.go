package service

import (
	"errors"
	"io/ioutil"

	"github.com/JabinGP/ssr-sub-hub/config"
	"github.com/JabinGP/ssr-sub-hub/model/pojo"
)

// UserService user serice
type UserService struct {
}

// Verify verify username and password
func (userService *UserService) Verify(username string, password string) error {
	var flagExistName bool
	config.FlashConfig()
	userList := []pojo.User{}
	config.Viper.UnmarshalKey("user", &userList)

	for _, user := range userList {
		if user.Username == username {
			flagExistName = true
			if user.Password != password {
				return errors.New("用户名密码不匹配")
			}
			break
		}
	}

	if !flagExistName {
		return errors.New("用户不存在")
	}

	return nil
}

// GetSSRList Get User SSR List
func (userService *UserService) GetSSRList(username string) ([]pojo.SSR, error) {
	ssrList := []pojo.SSR{}

	userConfig, err := config.GetUserConfig(username)
	if err != nil {
		return nil, err
	}
	userConfig.UnmarshalKey("ssr", &ssrList)
	return ssrList, nil
}

// GetV2rayList Get User V2ray List
func (userService *UserService) GetV2rayList(username string) ([]pojo.V2ray, error) {
	v2rayList := []pojo.V2ray{}

	userConfig, err := config.GetUserConfig(username)
	if err != nil {
		return nil, err
	}
	userConfig.UnmarshalKey("v2ray", &v2rayList)
	return v2rayList, nil
}


// UpdateUserConfig ...
func (userService *UserService) UpdateUserConfig(username string, userConfig string) error {
	configByteList := []byte(userConfig)
	err := ioutil.WriteFile("./user/"+username+".toml",configByteList,0666)
	if err != nil {
		return err
	}
	return nil
}
