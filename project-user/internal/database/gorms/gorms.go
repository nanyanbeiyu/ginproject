/*
@author: NanYan
*/
package gorms

import (
	"carrygpc.com/project-user/config"
	"carrygpc.com/project-user/internal/data/member"
	"carrygpc.com/project-user/internal/data/organization"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	//配置MySQL连接参数
	username := config.C.MC.Username //账号
	password := config.C.MC.Password //密码
	host := config.C.MC.Host         //数据库地址，可以是Ip或者域名
	port := config.C.MC.Port         //数据库端口
	Dbname := config.C.MC.Dbname     //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	_db.AutoMigrate(&member.Member{}, &organization.Organization{})
}

func GetDb() *gorm.DB {
	return _db
}

type GormConn struct {
	DB *gorm.DB
	TX *gorm.DB
}

func New() *GormConn {
	return &GormConn{
		DB: GetDb(),
		TX: GetDb(),
	}
}

func (g *GormConn) Default(ctx context.Context) *gorm.DB {
	return g.DB.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Rollback() {
	g.TX.Rollback()
}
func (g *GormConn) Commit() {
	g.TX.Commit()
}

func (g *GormConn) Tx(ctx context.Context) *gorm.DB {
	return g.TX.WithContext(ctx)
}

func (g *GormConn) Begin() {
	g.TX = GetDb().Begin()
}
