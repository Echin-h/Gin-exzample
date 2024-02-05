package errs

// 这个代码是用作特定类型定义
//可以预定义一些成功类型和错误类型，以便在整个系统中使用

//参照NX的博客，反正就是参照了HTTP状态码的语义，方便识别错误类型

// 这里的code都是*Error的类型
var (
	SUCCESS = newError(200, "Success")
)

var (
	INVALID_REQUEST = newError(40001, "无效的请求")
	NOTFOUND        = newError(40002, "目标不存在")
	HAS_EXIST       = newError(40003, "目标已存在")
	LOGIN_ERROR     = newError(40004, "d登陆失败")
	UNTHORIZATION   = newError(40005, "鉴权失败")
)

var (
	DB_LINK_ERROR = newError(50001, "连接数据库失败")
	DB_CRUD_ERROR = newError(50002, "数据库操作失败")
	DB_BASE_ERROR = newError(50003, "数据库内部错误")
)

var (
	SERVE_INTERNAL = newError(60001, "服务器内部故障")
)
