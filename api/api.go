package api

type Query interface{}

//请求结构体
type QueryParams struct {
	Qid string
	Query Query
}

type Error struct {
	ErrId int		`json:"error_id"`
	ErrStr string	`json:"error_str"`
	Msg string		`json:"msg"`
}

type Data interface{}

//相应结构体
type Response struct {
	Error Error	`json:"error"`
	Data Data	`json:"data"`
}

func NewError(err_id int, err_str string, msg string) *Error {
	return &Error{ErrId: err_id, ErrStr: err_str, Msg: msg}
}
