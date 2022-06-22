package apk

var cfg ConfigJson

type ConfigJson struct {
	ApkTool   ApkToolParams   `json:"apktool"`
	ApkSigner ApkSignerParams `json:"apksigner"`
}

type ApkToolParams struct {
	// apktool所在路径
	ApkToolPath string `json:"apktoolpath"`
	// 输出路径
	OutDir string `json:"outdir"`
}

type JarSignerParams struct {
	// java 路径
	JarSignerPath string
	// 密钥存放路径
	KeyStorePath string
	// 签名后的生成路径
	SignedJarPath string
	// 别名
	KeyAlias string
	// 密钥文件密码
	StorePassword string
	// 要签名的apk路径
	SourceApkPath string
}

type ApkSignerParams struct {
	// apksigner 路径,一般在sdk对应版本下
	ApkSignerPath string `json:"apksignerpath"`
	// 密钥存放路径
	KeyStorePath string `json:"ks"`
	// 签名后的生成路径
	OutPath string `json:"out"`
	// 别名
	KeyAlias string `json:"alias"`
	// 密钥文件密码
	StorePassword string `json:"kspwd"`
	// 要签名的apk路径
	InPath string `json:"in"`
}
