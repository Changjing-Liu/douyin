package common

import "errors"

var (
	ErrorUserNameNull    = errors.New("用户名为空")
	ErrorUserNameExtend  = errors.New("用户名长度不符合规范")
	ErrorPasswordNull    = errors.New("密码为空")
	ErrorPasswordLength  = errors.New("密码长度不符合规范")
	ErrorUserExit        = errors.New("用户已存在")
	ErrorUserNotExit     = errors.New("用户不存在")
	ErrorPasswordWrong   = errors.New("密码错误")
	ErrorFullPossibility = errors.New("用户不存在，账号或密码出错")
	ErrorGenIDFailed     = errors.New("创建用户ID失败")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorQueryFailed     = errors.New("查询数据失败")
	ErrorInsertFailed    = errors.New("插入数据失败")
	ErrorNullPointer     = errors.New("空指针异常")
	ErrEmptyUserList     = errors.New("用户列表为空")
)