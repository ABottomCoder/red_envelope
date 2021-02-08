package main

import (
	_ "github.com/red_envelope"
	"github.com/red_envelope/infra"
	"github.com/red_envelope/infra/base"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main()  {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()

}
