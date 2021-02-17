package errmsg

//错误码
const (
	Success = 200
	Error   = 500
	//请求类错误
	ErrRequest = 301
	//code1000开头的是用户错误
	ErrUserNameUsed        = 1001
	ErrUserPasswordWrong   = 1002
	ErrUserNotExist        = 1003
	ErrUserTokenNotExist   = 1004
	ErrUserNotHaveAddRight = 1005
	ErrUserEmailUsed       = 1006
	//code2000开头的是文章模块错误
	ErrArticleNotExist = 2001
	//code3000开头的是分类模块错误
	ErrCategoryUsed     = 3001
	ErrCategoryIdUsed   = 3002
	ErrCategoryNotExist = 3003
	//token类错误
	InvalidToken      = 4001
	TokenNotExist     = 4002
	TokenError        = 4003
	AuthEmpty         = 4004
	TokenRunTimeError = 4005
	//登录类错误
	ErrPassword     = 5001
	ErrNotHaveRight = 5002
	//数据库查找类错误
	ErrDatabaseFind = 6001
)

//错误信息字典
var CodeMsg = map[int]string{
	Success:                "ok",
	Error:                  "error",
	ErrUserNameUsed:        "用户名已经使用",
	ErrUserNotExist:        "用户名不存在",
	ErrUserPasswordWrong:   "用户密码错误",
	ErrUserTokenNotExist:   "token不存在",
	ErrCategoryUsed:        "目录名已经存在",
	ErrCategoryIdUsed:      "目录id已经存在",
	ErrCategoryNotExist:    "该目录不存在",
	InvalidToken:           "token不合法",
	TokenNotExist:          "token不存在",
	TokenError:             "token错误",
	AuthEmpty:              "请求头中auth为空",
	TokenRunTimeError:      "token过期",
	ErrRequest:             "请求错误",
	ErrPassword:            "密码错误",
	ErrNotHaveRight:        "用户无登录权限",
	ErrDatabaseFind:        "数据库查找错误",
	ErrUserNotHaveAddRight: "用户没有添加权限",
	ErrUserEmailUsed:       "用户邮箱被使用",
}

////获取错误码对应的信息
//func GetErrMsg(code int) string {
//	return CodeMsg[code]
//}
