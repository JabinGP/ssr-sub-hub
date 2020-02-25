package tool

import (
	"encoding/base64"
)

// DecodeBase64 解密base64编码
func DecodeBase64(src []byte) ([]byte, error) {
	encoder := getDecoder(src)
	n := encoder.DecodedLen(len(src)) //DecodedLen返回len(encoded)字节base64编码的数据解码后的最大长度
	encodedData := make([]byte, n)
	_, err := encoder.Decode(encodedData, src)
	if err != nil {
		return nil, err
	}
	return encodedData, nil
}

func getDecoder(src []byte) *base64.Encoding {
	var encoder = base64.RawStdEncoding
	if src[len(src)-1] == '=' {
		encoder = base64.StdEncoding
	}
	return encoder
}
