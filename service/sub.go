package service

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/JabinGP/ssr-sub-hub/config"
	"github.com/JabinGP/ssr-sub-hub/redis"
	"github.com/JabinGP/ssr-sub-hub/tool"
)

// SubService sub service
type SubService struct {
}

// GetAllSSR 下载所有连接的文件并集合到一个文件中
func (subService *SubService) GetAllSSR(linkList []string) ([]byte, error) {
	var dataList [][]byte
	var encodedDataList [][]byte
	var md5String = ""
	jsonBts := tool.GetJSONBts(linkList)
	if jsonBts == nil {
		return nil, errors.New("获取列表md5失败：转换json时出错")
	}
	md5String = fmt.Sprintf("%x", md5.Sum(jsonBts))

	if redis.Enable {
		log.Printf("Redis enabled, trying to get the value of key %s \n", md5String)
		res, err := redis.Client.Get(md5String).Result()
		if err == nil {
			log.Printf("Redis enabled, successfully got the value of key %s \n", md5String)
			return []byte(res), nil
		}
		log.Printf("Redis enabled, fail to get the value of key %s \n", md5String)
	}

	for _, url := range linkList {
		// 请求数据
		log.Printf("Downloading %s", url)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// 暂存数据
		dataList = append(dataList, data)
	}
	// 解密
	for index, data := range dataList {
		log.Printf("decoding No.%d", index+1)
		encodedData, err := tool.DecodeBase64(data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		encodedDataList = append(encodedDataList, encodedData)
	}

	// 拼接暂存数据
	allData := bytes.Join(encodedDataList, []byte{})

	// base64 编码拼接完成的数据
	n := base64.StdEncoding.EncodedLen(len(allData))
	decodedData := make([]byte, n)
	base64.StdEncoding.Encode(decodedData, allData)

	if redis.Enable {
		log.Printf("Redis enabled, trying to save the value of key %s \n", md5String)
		expiration := time.Duration(config.Viper.GetInt("redis.expiration"))
		set, err := redis.Client.SetNX(md5String, string(decodedData), expiration*time.Second).Result()
		if err != nil || set != true {
			log.Println(err)
		}
		log.Printf("Redis enabled, successfully saved the value of key %s \n", md5String)
	}
	return decodedData, nil
}
