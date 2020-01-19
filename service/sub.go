package service

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

// SubService sub service
type SubService struct {
}

// GetAllSSR 下载所有连接的文件并集合到一个文件中
func (subService *SubService) GetAllSSR(linkList []string) ([]byte, error) {
	var dataList [][]byte
	var encodedDataList [][]byte

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

		n := base64.RawStdEncoding.DecodedLen(len(data)) //DecodedLen返回len(encoded)字节base64编码的数据解码后的最大长度
		encodedData := make([]byte, n)
		_, err := base64.RawStdEncoding.Decode(encodedData, data)
		if err != nil {
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

	return decodedData, nil
}
