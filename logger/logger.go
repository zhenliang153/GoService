package logger

import (
	"os"
	"fmt"
	"time"
	"runtime"
	"strings"
	"github.com/sirupsen/logrus"
)

var rlog = logrus.New()
var spErrorLogger *logrus.Entry

// Msg of log fields
type Msg logrus.Fields

// GetLogger return default rlog
func GetLogger() *logrus.Logger {
	return rlog
}

// GetSpErrorLogger return sp_error_rlog
func GetSpErrorLogger() *logrus.Entry {
	return spErrorLogger
}

func init() {
	rlog.SetFormatter(&logrus.TextFormatter{})
	rlog.SetOutput(os.Stdout)
	rlog.SetLevel(logrus.TraceLevel)

	spErrorLogger = rlog.WithFields(logrus.Fields{"_type": "sp_error"})
}

//获取文件名称
func getPathStr() string{
	_, file, line, _ := runtime.Caller(3)
	//return fmt.Sprintf("%s:%d", file[strings.Index(file, "/Filter/"):], line)
	return fmt.Sprintf("%s:%d", file[strings.LastIndex(file, "/")+1:], line)
}
//获取时间，now.Month()等方法获取个位数时前面没"0"
func getTimeStr() string {
	now := time.Now()
	//fmt.Println(now.Format("2006/01/02 15:04:05"))
	//return fmt.Sprintf("%d%d%d %d:%d:%d", now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	return fmt.Sprintf(now.Format("20060102 15:04:05"))
}
//通过循环添加可变参数
func log_base(tt string, ss string, params ...interface{}){
	//直接输出会报错 %!(EXTRA type=value)
	//fmt.Printf("[%s][%s][%s]%s\n", tt, getTimeStr(), getPathStr(), fmt.Sprintf(ss, params...) )
	str := ss
	for _, param := range params {
		str += fmt.Sprintf("%v", param)
	}
	fmt.Printf("[%s][%s][%s]%s\n", tt, getTimeStr(), getPathStr(), str)
}

//建议调用以下方法
func LOG_DEBUG(ss string, params ...interface{}){
	log_base("DEBUG", ss, params...)
}
func LOG_INFO(ss string, params ...interface{}){
	log_base("INFO", ss, params...)
}
func LOG_ERROR(ss string, params ...interface{}){
	log_base("ERROR", ss, params...)
}
