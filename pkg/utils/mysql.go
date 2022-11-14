package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/dnsjia/luban/cmd/options"
)

func GormMySQL() *gorm.DB {

	m := options.Config.Mysql

	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Config

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	var err error
	if options.DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	}); err != nil {
		panic(fmt.Errorf("无法连接到数据库"))
	}

	return options.DB
}
