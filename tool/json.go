package tool

import "encoding/json"

// GetJSONString ...
func GetJSONString(src interface{}) string {
	jsonBts := GetJSONBts(src)
	if jsonBts == nil {
		return ""
	}
	return string(jsonBts)
}

// GetJSONBts ...
func GetJSONBts(src interface{}) []byte {
	jsonBts, err := json.Marshal(src)
	if err != nil {
		return nil
	}
	return jsonBts
}
