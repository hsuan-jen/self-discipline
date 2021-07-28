package response

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	// 服务级错误码
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	CallHTTPError      = 10105
	ResubmitError      = 10106
	ResubmitMsg        = 10107
	HashIdsDecodeError = 10108
	SignatureError     = 10109
	RBACError          = 10110

	// 业务模块级错误码
	// 用户模块
	IllegalUserName      = 20101
	UserCreateError      = 20102
	UserUpdateError      = 20103
	UserSearchError      = 20104
	UserSetStatusErr     = 20105
	UserRecordErr        = 20106
	UserEmpty            = 20107
	WechatAccessTokenErr = 20108
	WechatUserInfoErr    = 20109

	// 配置
	ConfigEmailError        = 20401
	ConfigSaveError         = 20402
	ConfigRedisConnectError = 20403
	ConfigMySQLConnectError = 20404
	ConfigMySQLInstallError = 20405
	ConfigGoVersionError    = 20406

	// 实用工具箱
	SearchRedisError = 20501
	ClearRedisError  = 20502
	SearchRedisEmpty = 20503
	SearchMySQLError = 20504

	JwtError       = 20601
	JwtOverdue     = 20602
	JwtResoluErr   = 20603
	JwtCreateError = 20604
)

var codeText = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败",
	ResubmitError:      "Resubmit Error",
	ResubmitMsg:        "请勿重复提交",
	HashIdsDecodeError: "ID 参数有误",
	SignatureError:     "Signature Error",

	IllegalUserName:      "非法用户名",
	UserCreateError:      "创建用户失败",
	UserUpdateError:      "更新用户失败",
	UserSearchError:      "查询用户失败",
	UserSetStatusErr:     "设置登录状态失败",
	UserRecordErr:        "记录登录信息失败",
	UserEmpty:            "用户名不存在或者密码错误",
	WechatAccessTokenErr: "获取access_token失败",
	WechatUserInfoErr:    "获取微信用户信息失败",

	JwtError:       "未登录或非法访问",
	JwtOverdue:     "授权已过期",
	JwtResoluErr:   "jwt解析错误",
	JwtCreateError: "获取token失败",

	ConfigEmailError:        "修改邮箱配置失败",
	ConfigSaveError:         "写入配置文件失败",
	ConfigRedisConnectError: "Redis 连接失败",
	ConfigMySQLConnectError: "MySQL 连接失败",
	ConfigMySQLInstallError: "MySQL 初始化数据失败",

	SearchRedisError: "查询 Redis Key 失败",
	ClearRedisError:  "清空 Redis Key 失败",
	SearchRedisEmpty: "查询的 Redis Key 不存在",
	SearchMySQLError: "查询 MySQL 失败",
}

func Text(code int) string {
	return codeText[code]
}
