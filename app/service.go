package app

import (
	"Service/api"
	//"Service/init"
	"Service/logger"
)

func serviceFunc(req *api.QueryParams, resp *api.Response) {
	query := req.Query.(*ServiceReq)
	mode := query.Mode
	text_in := query.Text
	logger.LOG_INFO("Mode: ", mode, " Text: ", text_in)
	var text_out string
	//logger.LOG_INFO("text_out: ", text_out)	
	var resp_data ServiceResp
	resp_data.TextIn = text_in
	resp_data.TextOut = text_out
	resp.Data = resp_data
	var resp_err api.Error
	resp_err.ErrId = 0
	resp_err.ErrStr = "success"
	resp_err.Msg = "success"
	resp.Error = resp_err
}

//请求的数据结构
type ServiceReq struct {
	Mode int
	Text string
}

//返回的数据结构
type ServiceResp struct {
	TextIn string	`json:"text_in"`
	TextOut string	`json:"text_out"`
}
