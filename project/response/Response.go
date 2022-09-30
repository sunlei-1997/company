package response

import (
	"project/entity"
)

type (
	GlobalErrorCode struct {
		ErrorCode int    `json:"code"`
		Error     string `json:"msg"`
	}

	Response struct {
		MoreInfo  string      `json:"moreInfo"`
		Data      interface{} `json:"data"`
		ErrorCode int         `json:"code"`
		Error     string      `json:"msg"`
	}
	ProjectResponse struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	NewData struct {
		List  interface{} `json:"list"`
		Total int         `json:"total"`
	}
)

func OK() *Response {
	res := &Response{
		MoreInfo:  "",
		Data:      true,
		Error:     "",
		ErrorCode: 200,
	}
	return res
}

func Failure() *Response {
	res := &Response{
		MoreInfo:  "",
		Data:      false,
		Error:     "",
		ErrorCode: 200,
	}
	return res
}

// 查询所有
func All(data interface{}) *ProjectResponse {
	total := len(data.([]*entity.Food))
	newData := &NewData{List: data, Total: total}
	res := &ProjectResponse{
		Code: 1,
		Data: newData,
		Msg:  "查询成功",
	}
	return res
}

// 查询单个
func Single(data interface{}) *ProjectResponse {
	res := &ProjectResponse{
		Code: 1,
		Data: data,
		Msg:  "查询成功",
	}
	return res
}

func ErrorMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  msg,
		Error:     msg,
		ErrorCode: 500,
	}
	return res
}

func ParticularMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  msg,
		Error:     msg,
		ErrorCode: 501,
	}
	return res
}

func SuccessMsg(msg string) *Response {
	res := &Response{
		MoreInfo:  "",
		Error:     msg,
		ErrorCode: 200,
	}
	return res
}

func Error(g *GlobalErrorCode) *Response {
	if g == nil {
		g = &GlobalErrorCode{
			ErrorCode: 500,
			Error:     "",
		}
	}

	res := &Response{
		MoreInfo:  g.Error,
		Error:     g.Error,
		ErrorCode: g.ErrorCode,
	}
	return res
}

var (
	ParamError       = GlobalErrorCode{ErrorCode: 500, Error: "参数错误"}
	ServerUknowError = GlobalErrorCode{ErrorCode: 500, Error: "服务器发生未知异常"}
)
