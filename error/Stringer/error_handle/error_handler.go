package error_handle

import (
	"github.com/pkg/errors"
)

// 1、自定义error结构体，并重写Error()方法
// 错误时返回自定义结构
type CustomError struct {
	Code    ErrCode `json:"code"`    // 业务码
	Message string  `json:"message"` // 业务码
}

func (e *CustomError) Error() string {
	return e.Code.Str()
}

type ErrCode int64 //错误码

// 2、定义errorCode
//go:generate stringer -type ErrCode -linecomment
const (
	// 服务级错误码
	ServerError        ErrCode = 10101 // Internal Server Error
	TooManyRequests    ErrCode = 10102 // Too Many Requests
	ParamBindError     ErrCode = 10103 // 参数信息有误
	AuthorizationError ErrCode = 10104 // 签名信息有误
	CallHTTPError      ErrCode = 10105 // 调用第三方HTTP接口失败
	ResubmitError      ErrCode = 10106 // ResubmitError
	ResubmitMsg        ErrCode = 10107 // 请勿重复提交
	HashIdsDecodeError ErrCode = 10108 // ID参数有误
	SignatureError     ErrCode = 10109 // SignatureError

	// 业务模块级错误码
	// 用户模块
	IllegalUserName ErrCode = 20101 // 非法用户名
	UserCreateError ErrCode = 20102 // 创建用户失败
	UserUpdateError ErrCode = 20103 // 更新用户失败
	UserSearchError ErrCode = 20104 // 查询用户失败

	// 配置
	ConfigEmailError        ErrCode = 20401 // 修改邮箱配置失败
	ConfigSaveError         ErrCode = 20402 // 写入配置文件失败
	ConfigRedisConnectError ErrCode = 20403 // Redis连接失败
	ConfigMySQLConnectError ErrCode = 20404 // MySQL连接失败
	ConfigMySQLInstallError ErrCode = 20405 // MySQL初始化数据失败
	ConfigGoVersionError    ErrCode = 20406 // GoVersion不满足要求

	// 实用工具箱
	SearchRedisError ErrCode = 20501 // 查询RedisKey失败
	ClearRedisError  ErrCode = 20502 // 清空RedisKey失败
	SearchRedisEmpty ErrCode = 20503 // 查询的RedisKey不存在
	SearchMySQLError ErrCode = 20504 // 查询mysql失败

	// 菜单栏
	MenuCreateError ErrCode = 20601 // 创建菜单失败
	MenuUpdateError ErrCode = 20602 // 更新菜单失败
	MenuListError   ErrCode = 20603 // 删除菜单失败
	MenuDeleteError ErrCode = 20604 // 获取菜单列表页失败
	MenuDetailError ErrCode = 20605 // 获取菜单详情失败

	// 借书
	BookNotFoundError        ErrCode = 20701 // 书未找到
	BookHasBeenBorrowedError ErrCode = 20702 // 书已经被借走了
)

// 4、新建自定义error实例化
func NewCustomError(code ErrCode) error {
	// 初次调用得用Wrap方法，进行实例化
	return errors.Wrap(&CustomError{
		Code:    code,
		Message: code.Str(),
	}, "")
}
