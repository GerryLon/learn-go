package util

import "encoding/json"

// json response
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
}

// 将需要返回的数据, 以Response的格式转成json
func ToJson(data interface{}, ret int, msg string) (string, error) {

	b, err := json.Marshal(&Response{
		Data: data,
		Ret:  ret,
		Msg:  msg,
	})

	if err == nil {
		return "", err
	}
	return string(b), nil
}

func SuccessJson(data interface{}) string {
	json, err := ToJson(data, 0, "ok")
	if err != nil {
		panic(err)
		return ""
	}
	return json
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Data: data,
		Ret:  0,
		Msg:  "ok",
	}
}

func FailResponse(data interface{}) Response {
	return Response{
		Data: data,
		Ret:  -1,
		Msg:  "error",
	}
}
