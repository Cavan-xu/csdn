package main

import (
	"fmt"
	"time"
	"xorm.io/core"

	"github.com/go-xorm/xorm"
)

type MysqlConfig struct {
	Username    string
	Password    string
	Host        string
	Port        int
	DBName      string
	MaxIdleConn int
	MaxOpenConn int
	MaxAlive    int
}

type MysqlHandler struct {
	engine *xorm.Engine
}

var (
	mysqlHandler *MysqlHandler
)

func InitMysqlHandler(conf MysqlConfig) error {
	var (
		err error
	)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return err
	}

	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	engine.SetMaxOpenConns(conf.MaxOpenConn)    // 设置连接池最大打开连接数
	engine.SetMaxIdleConns(conf.MaxIdleConn)    // 设置连接池最大空闲连接数
	engine.SetConnMaxLifetime(10 * time.Second) // 设置连接最大复用时间
	engine.SetMapper(core.LintGonicMapper)

	err = engine.Ping()
	if err != nil {
		return err
	}

	mysqlHandler = &MysqlHandler{
		engine: engine,
	}
	return nil
}

func NewSession() *xorm.Session {
	return mysqlHandler.engine.NewSession()
}
