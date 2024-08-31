package dao

import "gorm.io/gorm"

var DB *gorm.DB

func Create[T any](data *T) {
	err := DB.Create(&data).Error
	if err != nil {
		panic(HandleMySQLError(err))
	}
}

// 查询获取一条数据
func QueryRow[T any](row T, query string, args ...any) T {
	err := DB.Where(query, args).First(&row).Error
	if err != nil {
		panic(HandleMySQLError(err))
	}
	return row
}
