/*
@author: NanYan
*/
package config

import (
	"carrygpc.com/project-common/logs"
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
}

type ServeConf struct {
	Addr string
	Name string
}

type GrpcConf struct {
	Addr string
	Name string
}

type EtcdConf struct {
	Addrs []string
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
	conf.GetEtcdConf()
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
		Addr: c.viper.GetString("serve.addr"),
		Name: c.viper.GetString("serve.name"),
	}
}

func (c *Conf) GetEtcdConf() {
	ec := &EtcdConf{
		Addrs: c.viper.GetStringSlice("etcd.addrs"),
	}
	c.EC = ec
}
