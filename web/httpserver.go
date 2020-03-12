package web

import (
	"Service/api"
	"Service/app"
	"Service/logger"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var allHTTPApi app.KeyMap

func init() {
	allHTTPApi = app.HTTPApis()
}

func getGetReqPara(context *gin.Context,req *api.QueryParams) {
	var ok bool
	req.Qid, ok = context.GetQuery("qid")
	if !ok {
		logger.LOG_ERROR("it has no qid!")
	}
}

//过滤逻辑函数体
func serviceQuery(c *gin.Context) {
	var req api.QueryParams
	var resp api.Response

	//初始化get请求参数
	getGetReqPara(c, &req)
	qtype, ok := c.GetQuery("type")
	if !ok {
		logger.LOG_ERROR("it has no type!")
		resp.Error = *api.NewError(0, "missing variables", "it has no type")
		c.JSON(http.StatusOK, resp)
		return
	}
	query, ok := c.GetQuery("query")
	if !ok {
		logger.LOG_ERROR("it has no query!")
		resp.Error = *api.NewError(0, "missing variables", "it has no query")
		c.JSON(http.StatusOK, resp)
		return
	}

	logger.LOG_INFO("type:", qtype, "query:", query)

	//判断是否正确接口
	apiInterface, isType := allHTTPApi[qtype]
	if !isType {
		logger.LOG_ERROR("Error type: ", qtype)
		resp.Error = *api.NewError(0, "error type", "it has no type in apis")
		c.JSON(http.StatusOK, resp)
		return
	}

	//获取当前接口的相关接口
	currentAPI := apiInterface.(app.KeyMap)
	apiFunction := currentAPI["fun"].(func(*api.QueryParams,*api.Response))
	apiQueryGen := currentAPI["req"].(func() api.Query)
	apiQuery := apiQueryGen()
	if apiQuery == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	if reqErr := json.Unmarshal([]byte(query), &apiQuery); reqErr != nil {
		logger.LOG_ERROR("convert query to object error")
		resp.Error = *api.NewError(0, "format err", "convert query to object error")
		c.JSON(http.StatusOK, resp)
		return
	}
	req.Query = apiQuery
	//监控超时事件，有需要再添加
	//设置结束通知chan
	var finish_chan = make(chan int)
	go func(req *api.QueryParams, resp *api.Response) {
		apiFunction(req, resp)
		finish_chan <- 0
	}(&req, &resp)
	value := <- finish_chan
	logger.LOG_INFO("value: ", value)
	c.JSON(http.StatusOK, resp)
}

//Start the WebService
func Start(port string) {
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", serviceQuery)
	//router.POST("/", serviceQuery)

	router.Run(":"+string(port))
}
