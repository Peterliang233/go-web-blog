package errmsg

// 错误码
const (
	Success = 200
	Error   = 500
	// 请求类错误
	ErrRequest = 301
	// code1000开头的是用户错误
	ErrUserNameUsed        = 1001
	ErrUserPasswordWrong   = 1002
	ErrUserNotExist        = 1003
	ErrUserTokenNotExist   = 1004
	ErrUserNotHaveAddRight = 1005
	ErrUserEmailUsed       = 1006
	ErrEmailUnVerify       = 1007
	// code2000开头的是文章模块错误
	ErrArticleNotExist = 2001
	// code3000开头的是分类模块错误
	ErrCategoryUsed     = 3001
	ErrCategoryIDUsed   = 3002
	ErrCategoryNotExist = 3003
	// token类错误
	InvalidToken      = 4001
	TokenNotExist     = 4002
	TokenError        = 4003
	AuthEmpty         = 4004
	TokenRunTimeError = 4005
	// 登录类错误
	ErrPassword     = 5001
	ErrNotHaveRight = 5002
	// 数据库查找类错误
	ErrDatabaseNotFound = 6001
	ErrDatabaseCreate   = 6002
	// 评论类的错误
	ErrCommentCreate = 7001
	ErrCommentDelete = 7002
	ErrCommentGet    = 7003
	// 标签类错误
	ErrTagCreate = 8001
	ErrTagGet    = 8002
	ErrTagDelete = 8003
	// 点赞类错误
	ErrLikeCreate = 9001
	ErrLikeGet    = 9002
)

// CodeMsg 错误信息字典
var CodeMsg = map[int]string{
	Success:                "ok",
	Error:                  "error",
	ErrUserNameUsed:        "用户名已经使用",
	ErrUserNotExist:        "用户名不存在",
	ErrUserPasswordWrong:   "用户密码错误",
	ErrUserTokenNotExist:   "token不存在",
	ErrCategoryUsed:        "目录名已经存在",
	ErrCategoryIDUsed:      "目录id已经存在",
	ErrCategoryNotExist:    "该目录不存在",
	InvalidToken:           "token不合法",
	TokenNotExist:          "token不存在",
	TokenError:             "token错误",
	AuthEmpty:              "请求头中auth为空",
	TokenRunTimeError:      "token过期",
	ErrRequest:             "请求错误",
	ErrPassword:            "密码错误",
	ErrNotHaveRight:        "用户无登录权限",
	ErrDatabaseNotFound:    "数据库查找错误",
	ErrUserNotHaveAddRight: "用户没有添加权限",
	ErrUserEmailUsed:       "用户邮箱被使用",
	ErrEmailUnVerify:       "邮箱未被激活",
	ErrCommentCreate:       "创建评论错误",
	ErrCommentDelete:       "删除评论错误",
	ErrCommentGet:          "获取评论错误",
	ErrTagCreate:           "创建标签错误",
	ErrTagGet:              "获取标签错误",
	ErrTagDelete:           "删除标签错误",
	ErrLikeCreate:          "创建点赞错误",
	ErrLikeGet:             "获取点赞错误",
}
