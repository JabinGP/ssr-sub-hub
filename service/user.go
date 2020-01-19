package service

import (
	"errors"
	"fmt"

	"github.com/JabinGP/ssr-sub-hub/config"
)

// UserService user serice
type UserService struct {
}

// Verify verify username and password
func (userService *UserService) Verify(username string, password string) error {
	config.FlashConfig()
	key := config.Viper.GetString(fmt.Sprintf("user.%s.key", username))

	if key == "" {
		return errors.New("用户不存在")
	}

	if key != password {
		return errors.New("用户名密码不匹配")
	}

	return nil
}

// GetSSRLinkList get user link list
func (userService *UserService) GetSSRLinkList(username string) []string {
	config.FlashConfig()
	linkList := config.Viper.GetStringSlice(fmt.Sprintf("user.%s.ssr", username))
	return linkList
}
