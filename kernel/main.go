package main

import (
	"fmt"
	"kernel/pkg/apk"
	"kernel/pkg/common"
)

const (
	id_apktool_unpack int = 1
	id_apktool_pack   int = 4
	id_add_log        int = 6
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
	fmt.Println("2. baksmali")
	fmt.Println("3. smali")
	fmt.Println("4. apktool给apk合包")
	fmt.Println("5. 给APK签名")
	fmt.Println("6. 增加日志模块代码")
	fmt.Println("0. 退出")
}

func routeMainMenu(id int) {
	switch id {
	case id_apktool_unpack:
		apk.Unpack()
	case id_apktool_pack:
		apk.Pack()
	case 2:
		apk.ToSign()
	case id_add_log:
	}
}
