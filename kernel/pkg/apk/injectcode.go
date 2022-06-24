package apk

import "kernel/pkg/common"

func InjectLogCode() {
	// 确定工程所在目录
	var to string = "/Users/koifishly/Desktop/cnfix/jebfenxi/"

	// smali文件
	var from string = "./cfgs/injectcode/Log/smali"

	// 复制文件
	common.CopyFile(from, to)
}
