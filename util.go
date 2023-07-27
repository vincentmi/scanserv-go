package main

import (
	"math/rand"
	"time"
	"unsafe"
)

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

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func randStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
