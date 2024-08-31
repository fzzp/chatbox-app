package dao

import (
	"chatbox-app/lib/errs"
	"errors"

	"gorm.io/gorm"
)

// type DBError struct {
// 	Err error
// }

// func NewDBError(err error) *DBError {
// 	return &DBError{err}
// }

// func (dbe *DBError) Error() string {
// 	if dbe.Err != nil {
// 		return dbe.Err.Error()
// 	}
// 	return ""
// }

// func (dbe *DBError) Unwrap() error {
// 	return dbe.Err
// }

// HandleMySQLError 处理MySQL错误，返回 *ApiError 和是否能处理
//
// 如果不是MyQL错误一律返回 *ErrServerError
func HandleMySQLError(err error) *errs.ApiError {
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
		return errs.ErrNotFound.AsException(err)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.ErrRecordExists.AsException(err)
	}

	return errs.ErrServerError.AsException(err)
}
