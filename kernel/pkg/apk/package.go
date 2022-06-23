package apk

import (
	"fmt"
	"os/exec"
)

func Unpack() {
	cmd := exec.Command(
		"java", "-jar", cfg.ApkTool.ApkToolPath,
		"d", cfg.ApkTool.InAPK,
		"-o", cfg.ApkTool.OutDir,
		"-f")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	fmt.Println(string(out))
}

func Pack() {
	cmd := exec.Command(
		"java", "-jar", "/Users/koifishly/bin/apktool_2.6.0.jar",
		"b", "/Users/koifishly/Desktop/cnfix/jebfenxi/out",
		"-o", "/Users/koifishly/Desktop/cnfix/jebfenxi/new.apk")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	fmt.Println(string(out))
}
