package main

import (
	"log"
	"syncAgent-go/syncAgent/httprequest"
)

func init() {
	//初始化CSRF + deviceID
	err := httprequest.CSRF()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//初始化globleConfig
	err = httprequest.GlobleConfigLoad()
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func main() {
	/*---------------------------------------------
	初始化
	1、获取全局config
	2、获取本地deviceID
	3、获取CSRF
	----------------------------------------------*/
	//启动QUIC 服务器等待接收指令

	//解析指令集

	//switch 指令调用相关API

	//返回结果
}
