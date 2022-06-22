package main

import (
	"fmt"
	"kernel/pkg/apk"
	"kernel/pkg/common"
)

const (
	id_apktool int = 1
)

func main() {
	apk.InitConfig()
	for {
		printMainMenu()
		id := common.ScanInt("选择功能：")
		if id == 0 {
			break
		}
		routeMainMenu(id)
	}

}

func printMainMenu() {
	fmt.Println("1. apktool给apk解包")
	fmt.Println("2. 给APK签名")
	fmt.Println("0. 退出")
}

func routeMainMenu(id int) {
	switch id {
	case id_apktool:
	case 2:
		apk.ToSign()
	}
}
