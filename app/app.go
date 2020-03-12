package app

import "Service/api"

type KeyMap map[string]interface{}

//API接口注册
var apis = KeyMap {
	"service": KeyMap {
		//主流程
		"fun": func(req *api.QueryParams,resp *api.Response) {
			serviceFunc(req, resp)
		},
		//返回请求的空数据结构
		"req": func() api.Query {
			query := &ServiceReq{}
			return query
		},
	},
}

func HTTPApis() KeyMap {
	return apis
}
