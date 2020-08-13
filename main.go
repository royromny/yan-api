package main

import (
	"yan_api/conf"
	_ "yan_api/conf"
	"yan_api/server"
	"yan_api/util"
)

func main() {
	//fmt.Printf("%+v\n", conf.Data)
	util.Log().Info("main %+v\n", conf.Data)

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
