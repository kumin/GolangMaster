package infras

import (
	"time"

	"gorm.io/gorm"
)

type MysqlConnector struct {
	db *gorm.DB
}

func MysqlConntionBuilder(fns ...optionFn) *MysqlConnector {
	opt := MysqlDefaultOption
	for _, f := range fns {
		f(opt)
	}
	return &MysqlConnector{}
}

func NewMysqlConnector() *MysqlConnector {
	return MysqlConnectionBuilder(
		WithDSN("root:root@tcp(localhost:3306)/tiki_anti_fraud_dev?charset=utf8&parseTime=True&loc=Local&multiStatements=true"),
	)
}

type MysqlOption struct {
	DSN         string
	MaxConn     int
	MaxLifeTime time.Duration
}

var MysqlDefaultOption = &MysqlOption{
	MaxConn:     2,
	MaxLifeTime: 1 * time.Minute,
}

type optionFn func(opt *MysqlOption)

func WithDSN(dsn string) optionFn {
	return func(opt *MysqlOption) {
		opt.DSN = dsn
	}
}

func WithMaxConn(conns int) optionFn {
	return func(opt *MysqlOption) {
		opt.MaxConn = conns
	}
}

func WithLifeTime(minus time.Duration) optionFn {
	return func(opt *MysqlOption) {
		opt.MaxLifeTime = minus
	}
}
