package main

import (
	"fmt"
	"yan_api/conf"
	_ "yan_api/conf"
	"yan_api/server"
)

func main() {
	fmt.Printf("%+v\n", conf.Data)

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
