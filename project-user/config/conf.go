/*
@author: NanYan
*/
package config

import (
	"carrygpc.com/project-common/logs"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"os"
)

var C = InitConfig()

type Conf struct {
	viper     *viper.Viper
	ServeConf *ServeConf
	GC        *GrpcConf
	EC        *EtcdConf
	MC        *MysqlConf
}

type ServeConf struct {
	Port string
	Host string
}

type GrpcConf struct {
	Addr    string
	Name    string
	Version string
	Weight  int64
}

type EtcdConf struct {
	Addrs []string
}

type MysqlConf struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func InitConfig() *Conf {
	v := viper.New()
	conf := &Conf{
		viper: v,
	}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.InitZapLog()
	conf.ServeConf = conf.GetServeConf()
	conf.GC = conf.GetGrpcConf()
	conf.GetEtcdConf()
	conf.GetMysqlConf()
	return conf
}

func (c *Conf) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln("11111", err)
	}
}

func (c *Conf) GetServeConf() *ServeConf {
	return &ServeConf{
		Port: c.viper.GetString("serve.port"),
		Host: c.viper.GetString("serve.host"),
	}
}

func (c *Conf) InitRedisConf() *redis.Options {
	return &redis.Options{
		Password: c.viper.GetString("gredis.password"),
		Addr:     c.viper.GetString("gredis.addr"),
		DB:       c.viper.GetInt("gredis.db"),
	}
}

func (c *Conf) GetGrpcConf() *GrpcConf {
	return &GrpcConf{
		Addr:    c.viper.GetString("grpc.addr"),
		Name:    c.viper.GetString("grpc.name"),
		Version: c.viper.GetString("grpc.version"),
		Weight:  c.viper.GetInt64("grpc.weight"),
	}
}

func (c *Conf) GetEtcdConf() {
	ec := &EtcdConf{
		Addrs: c.viper.GetStringSlice("etcd.addrs"),
	}
	c.EC = ec
}

func (c *Conf) GetMysqlConf() {
	mc := &MysqlConf{
		Username: c.viper.GetString("mysql.username"),
		Password: c.viper.GetString("mysql.password"),
		Host:     c.viper.GetString("mysql.host"),
		Port:     c.viper.GetString("mysql.port"),
		Dbname:   c.viper.GetString("mysql.dbname"),
	}
	c.MC = mc
}
