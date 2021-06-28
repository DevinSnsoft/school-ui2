package app

import "schoolServer/pkg/e"

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}

type Data struct {
	Token string `json:"token"`
}

type ResponseData struct {
	Roles []string `json:"roles"`
	Introduction string `json:"introduction"`
	Avatar string `json:"avatar"`
	Name string `json:"name"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = e.SUCCESS
	return res
}

func (res *Response) ReturnFalse() *Response {
	res.Code = e.ERROR
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

