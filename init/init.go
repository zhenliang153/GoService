package  initdata

import (
	"Service/logger"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Init() {
	//initMySQLData()
}

func initMySQLData() {
	//建立数据库连接池
	ip := viper.Get("connections.mysql.service.host").(string)
	port := viper.Get("connections.mysql.service.port").(int)
	database := viper.Get("connections.mysql.service.database").(string)
	user := viper.Get("connections.mysql.service.user").(string)
	passwd := viper.Get("connections.mysql.service.pass").(string)
	charset := viper.Get("connections.mysql.service.charset").(string)
	conn := getMySQLConn(ip,port,database,user,passwd,charset)
	defer conn.Close()
	initSensitiveWordTable(conn)

	logger.LOG_INFO("BadWord2State.size: ", len(BadWord2State))
}

func getMySQLConn(ip string,port int,database string,user string,passwd string,charset string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",user,passwd,ip,port,database,charset)
	conn, err := sql.Open("mysql",dsn)
	if err != nil{
		logger.LOG_ERROR("connect MySQL " + database + " failed!")
		panic("open database fail")
	}
	logger.LOG_INFO("connect MySQL " + database + " success!")
	return conn
}

func initSensitiveWordTable(conn *sql.DB) {
	sql := "SELECT bad_word,state FROM sensitive_word;"
	rows, err := conn.Query(sql)
	if err != nil {
		logger.LOG_ERROR("init sensitive_word error!")
		return
	}
	var bad_word string
	var state int
	for rows.Next() {
		rows.Scan(&bad_word, &state)
		BadWord2State[bad_word] = state
	}
}
