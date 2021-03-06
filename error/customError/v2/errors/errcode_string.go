// Code generated by "stringer -type errCode -linecomment"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerError-10101]
	_ = x[TooManyRequests-10102]
	_ = x[ParamBindError-10103]
	_ = x[AuthorizationError-10104]
	_ = x[CallHTTPError-10105]
	_ = x[ResubmitError-10106]
	_ = x[ResubmitMsg-10107]
	_ = x[HashIdsDecodeError-10108]
	_ = x[SignatureError-10109]
	_ = x[IllegalUserName-20101]
	_ = x[UserCreateError-20102]
	_ = x[UserUpdateError-20103]
	_ = x[UserSearchError-20104]
	_ = x[ConfigEmailError-20401]
	_ = x[ConfigSaveError-20402]
	_ = x[ConfigRedisConnectError-20403]
	_ = x[ConfigMySQLConnectError-20404]
	_ = x[ConfigMySQLInstallError-20405]
	_ = x[ConfigGoVersionError-20406]
	_ = x[SearchRedisError-20501]
	_ = x[ClearRedisError-20502]
	_ = x[SearchRedisEmpty-20503]
	_ = x[SearchMySQLError-20504]
	_ = x[MySQLNoQueryError-20505]
	_ = x[MySQLNoFieldError-20506]
	_ = x[MenuCreateError-20601]
	_ = x[MenuUpdateError-20602]
	_ = x[MenuListError-20603]
	_ = x[MenuDeleteError-20604]
	_ = x[MenuDetailError-20605]
	_ = x[BookNotFoundError-20701]
	_ = x[BookHasBeenBorrowedError-20702]
}

const (
	_errCode_name_0 = "Internal Server ErrorToo Many Requests参数信息有误签名信息有误调用第三方HTTP接口失败ResubmitError请勿重复提交ID参数有误SignatureError"
	_errCode_name_1 = "非法用户名创建用户失败更新用户失败查询用户失败"
	_errCode_name_2 = "修改邮箱配置失败写入配置文件失败Redis连接失败MySQL连接失败MySQL初始化数据失败GoVersion不满足要求"
	_errCode_name_3 = "查询RedisKey失败清空RedisKey失败查询的RedisKey不存在查询mysql失败查询mysql无结果查询mysql字段不存在"
	_errCode_name_4 = "创建菜单失败更新菜单失败删除菜单失败获取菜单列表页失败获取菜单详情失败"
	_errCode_name_5 = "书未找到书已经被借走了"
)

var (
	_errCode_index_0 = [...]uint8{0, 21, 38, 56, 74, 105, 118, 136, 150, 164}
	_errCode_index_1 = [...]uint8{0, 15, 33, 51, 69}
	_errCode_index_2 = [...]uint8{0, 24, 48, 65, 82, 108, 132}
	_errCode_index_3 = [...]uint8{0, 20, 40, 66, 83, 103, 129}
	_errCode_index_4 = [...]uint8{0, 18, 36, 54, 81, 105}
	_errCode_index_5 = [...]uint8{0, 12, 33}
)

func (i errCode) String() string {
	switch {
	case 10101 <= i && i <= 10109:
		i -= 10101
		return _errCode_name_0[_errCode_index_0[i]:_errCode_index_0[i+1]]
	case 20101 <= i && i <= 20104:
		i -= 20101
		return _errCode_name_1[_errCode_index_1[i]:_errCode_index_1[i+1]]
	case 20401 <= i && i <= 20406:
		i -= 20401
		return _errCode_name_2[_errCode_index_2[i]:_errCode_index_2[i+1]]
	case 20501 <= i && i <= 20506:
		i -= 20501
		return _errCode_name_3[_errCode_index_3[i]:_errCode_index_3[i+1]]
	case 20601 <= i && i <= 20605:
		i -= 20601
		return _errCode_name_4[_errCode_index_4[i]:_errCode_index_4[i+1]]
	case 20701 <= i && i <= 20702:
		i -= 20701
		return _errCode_name_5[_errCode_index_5[i]:_errCode_index_5[i+1]]
	default:
		return "errCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
