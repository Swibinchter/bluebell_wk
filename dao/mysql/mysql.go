package mysql

// 连接mysql数据库

import (
	"fmt"
	"goWebCli/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化连接mysql数据库
func Init(cfg *setting.MysqlConfig) (err error) {
	// 从配置中获取连接信息
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect mysql failed, err:%v\n", err)
		return
	}

	// 设置数据库连接数量
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return
}

func Close() {
	_ = db.Close()
}
