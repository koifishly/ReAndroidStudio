package apk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kernel/pkg/common"
	"os"
)

func InitConfig() {
	if common.PathExists("./config.json") {
		readFromFile()
	} else {
		f, err := os.Create("./config.json")
		if err != nil { //如果有错误 打印错误 返回
			fmt.Println("err=", err)
			return
		}
		fileContent, err := json.Marshal(cfg)
		if err != nil {
			ioutil.WriteFile("config.json", fileContent, 0666)
		}
		var jsonContent bytes.Buffer
		_ = json.Indent(&jsonContent, []byte(fileContent), "", "    ")

		f.WriteString(jsonContent.String())
		defer f.Close()
	}
}

func readFromFile() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
	fmt.Println(cfg.ApkSigner.KeyStorePath)
}

// func saveConfig2file() {
// 	fileContent, err := json.Marshal(GGlobalConfig)
// 	if err != nil {
// 		ioutil.WriteFile("config.json", fileContent, 0666)
// 	}
// }
