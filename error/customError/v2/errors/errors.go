package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// 1、自定义error结构体，并重写Error()方法
// 错误时返回自定义结构
type CustomError struct {
	Code errCode `json:"code"` // 业务码
	Err  error
}

func (e *CustomError) Error() string {
	errMsg := ""
	if e.Err != nil {
		errMsg = e.Err.Error()
	}
	return fmt.Sprintf("code: %d, description: %s, error: %s ", e.Code, e.Code.String(), errMsg)
}

type errCode int64 //错误码

// 定义errorCode
// 执行 go generate 生成 String 方法
//go:generate stringer -type errCode -linecomment
const (
	// 服务级错误码
	ServerError        errCode = 10101 // Internal Server Error
	TooManyRequests    errCode = 10102 // Too Many Requests
	ParamBindError     errCode = 10103 // 参数信息有误
	AuthorizationError errCode = 10104 // 签名信息有误
	CallHTTPError      errCode = 10105 // 调用第三方HTTP接口失败
	ResubmitError      errCode = 10106 // ResubmitError
	ResubmitMsg        errCode = 10107 // 请勿重复提交
	HashIdsDecodeError errCode = 10108 // ID参数有误
	SignatureError     errCode = 10109 // SignatureError

	// 业务模块级错误码
	// 用户模块
	IllegalUserName errCode = 20101 // 非法用户名
	UserCreateError errCode = 20102 // 创建用户失败
	UserUpdateError errCode = 20103 // 更新用户失败
	UserSearchError errCode = 20104 // 查询用户失败

	// 配置
	ConfigEmailError        errCode = 20401 // 修改邮箱配置失败
	ConfigSaveError         errCode = 20402 // 写入配置文件失败
	ConfigRedisConnectError errCode = 20403 // Redis连接失败
	ConfigMySQLConnectError errCode = 20404 // MySQL连接失败
	ConfigMySQLInstallError errCode = 20405 // MySQL初始化数据失败
	ConfigGoVersionError    errCode = 20406 // GoVersion不满足要求

	// 实用工具箱
	SearchRedisError  errCode = 20501 // 查询RedisKey失败
	ClearRedisError   errCode = 20502 // 清空RedisKey失败
	SearchRedisEmpty  errCode = 20503 // 查询的RedisKey不存在
	SearchMySQLError  errCode = 20504 // 查询mysql失败
	MySQLNoQueryError errCode = 20505 // 查询数据库无结果
	MySQLNoFieldError errCode = 20506 // 查询数据库字段不存在

	// 菜单栏
	MenuCreateError errCode = 20601 // 创建菜单失败
	MenuUpdateError errCode = 20602 // 更新菜单失败
	MenuListError   errCode = 20603 // 删除菜单失败
	MenuDeleteError errCode = 20604 // 获取菜单列表页失败
	MenuDetailError errCode = 20605 // 获取菜单详情失败

	// 借书
	BookNotFoundError        errCode = 20701 // 书未找到
	BookHasBeenBorrowedError errCode = 20702 // 书已经被借走了
)

func (e errCode) WrapWithMessage(err error, message string) error {
	return errors.Wrap(&CustomError{
		Code: e, Err: err,
	}, message)
}

func (e errCode) Wrap(err error) error {
	return errors.Wrap(&CustomError{
		Code: e, Err: err,
	}, "")
}
