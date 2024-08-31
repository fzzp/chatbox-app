package errs

import (
	"errors"

	"gorm.io/gorm"
)

// HandleMySQLError 处理MySQL错误，返回 *ApiError 和是否能处理
//
// 如果不是MyQL错误一律返回 *ErrServerError
func HandleMySQLError(err error) (*ApiError, bool) {
	// 非gorm定义错误处理方式
	// MySQL 错误码参照：
	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	// dbErr, ok := err.(*mysql.MySQLError)
	// if ok {
	// 	if dbErr.Number == 1062 {
	// 		return ErrRecordExists.AsException(err), true
	// 	}
	// }

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound.AsException(err), true
	}

	// gorm 必须将 TranslateError 设置为 true 才有效
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrRecordExists.AsException(err), true
	}

	return ErrServerError.AsException(err), false
}
