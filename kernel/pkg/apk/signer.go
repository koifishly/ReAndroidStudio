package apk

import (
	"fmt"
	"kernel/pkg/common"
	"os/exec"
)

var ApkSigner ApkSignerParams

func ToSign() {
	// 现在主要有两种签名方式。
	// JarSign 支持v1签名
	// ApkSign 支持v1、v2签名
	// 默认情况下都是采用ApkSign进行签名
	apksign()
}

func apksign() {
	fmt.Println("输入apksigner路径:")
	fmt.Println("当前设置为：" + ApkSigner.ApkSignerPath)
	var input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	if len(input) != 0 {
		ApkSigner.ApkSignerPath = input
	}

	// // 密钥存放路径
	// fmt.Println("签名文件所在路径:")
	// fmt.Println("当前设置为：" + ApkSigner.KeyStorePath)
	// input = ipt.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	ApkSigner.KeyStorePath = input
	// }

	// // 签名后的生成路径
	// // fmt.Println("签名后生成的路径:")
	// // fmt.Println("当前设置为：" + JarSigner.SignedJarPath)
	// // input = ipt.ScanString("Enter继续上次设置或者输入新的设置:")
	// // if len(input) != 0 {
	// // 	JarSigner.SignedJarPath = input
	// // }

	// // 别名
	// fmt.Println("别名:")
	// fmt.Println("当前设置为：" + ApkSigner.KeyAlias)
	// input = ipt.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	ApkSigner.KeyAlias = input
	// }

	// // 密钥文件密码
	// fmt.Println("密钥文件密码:")
	// fmt.Println("当前设置为：" + ApkSigner.StorePassword)
	// input = ipt.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	ApkSigner.StorePassword = input
	// }

	// // 要签名的apk路径
	// fmt.Println("要签名的apk路径:")
	// fmt.Println("当前设置为：" + ApkSigner.SourceApkPath)
	// input = ipt.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	ApkSigner.SourceApkPath = input
	// }

	apkSign()
}

func apkSign() {
	cmd := exec.Command(
		"/Users/koifishly/Library/Android/sdk/build-tools/30.0.3/apksigner",
		"sign",
		"--ks", "/Users/koifishly/Downloads/defaultApkSigner.keystore",
		"--ks-key-alias", "key0",
		"--ks-pass", "pass:123456",
		"/Users/koifishly/Downloads/1easy_java.apk")
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
	}
	fmt.Println(string(out))
}

var JarSigner JarSignerParams

func JarSignApk() {
	// jarsigner路径
	// fmt.Println("输入jarsigner路径:")
	// fmt.Println("当前设置为：" + JarSigner.JarSignerPath)
	// var input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.JarSignerPath = input
	// }

	// // 密钥存放路径
	// fmt.Println("签名文件所在路径:")
	// fmt.Println("当前设置为：" + JarSigner.KeyStorePath)
	// input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.KeyStorePath = input
	// }

	// // 签名后的生成路径
	// fmt.Println("签名后生成的路径:")
	// fmt.Println("当前设置为：" + JarSigner.SignedJarPath)
	// input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.SignedJarPath = input
	// }

	// // 别名
	// fmt.Println("别名:")
	// fmt.Println("当前设置为：" + JarSigner.KeyAlias)
	// input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.KeyAlias = input
	// }

	// // 密钥文件密码
	// fmt.Println("密钥文件密码:")
	// fmt.Println("当前设置为：" + JarSigner.StorePassword)
	// input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.StorePassword = input
	// }

	// // 要签名的apk路径
	// fmt.Println("要签名的apk路径:")
	// fmt.Println("当前设置为：" + JarSigner.SourceApkPath)
	// input = common.ScanString("Enter继续上次设置或者输入新的设置:")
	// if len(input) != 0 {
	// 	JarSigner.SourceApkPath = input
	// }

	javaSign()
}

func javaSign() {
	// jarsigner
	// -verbose
	// -keystore /Users/koifishly/code_projects/ReAndroidStudio/kernel/testkey
	// -signedjar /Users/koifishly/code_projects/ReAndroidStudio/kernel/test_signed.apk
	// /Users/koifishly/code_projects/ReAndroidStudio/kernel/test.apk
	// key0
	// -storepass 123456

	cmd := exec.Command(
		JarSigner.JarSignerPath,
		//		"-verbose",
		"-keystore", JarSigner.KeyStorePath,
		//"-signedjar", JarSigner.SignedJarPath,
		JarSigner.SourceApkPath,
		JarSigner.KeyAlias,
		"-storepass", JarSigner.StorePassword)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		//log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
