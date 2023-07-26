package main

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) Response {
	resp := Response{}
	resp.Code = 0
	resp.Message = "ok"
	resp.Data = data
	return resp
}

func Fail() Response {
	resp := Response{}
	resp.Code = 500
	resp.Message = "server fail"
	resp.Data = nil
	return resp
}

func FailMsg(message string) Response {
	resp := Response{}
	resp.Code = 500
	resp.Message = message
	resp.Data = nil
	return resp
}

func FailCodeMsg(code int, message string) Response {
	resp := Response{}
	resp.Code = code
	resp.Message = message
	resp.Data = nil
	return resp
}

func FailF(code int, message string, data interface{}) Response {
	resp := Response{}
	resp.Code = code
	resp.Message = message
	resp.Data = data
	return resp
}
