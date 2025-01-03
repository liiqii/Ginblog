package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisAddr     string
	RedisPort     string
	RedisUser     string
	RedisPassWord string

	Zone       int
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string

	Account string
	Pwd     string
	Url     string
)

// 初始化
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadRedis(file)
	LoadQiniu(file)
	LoadScan(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("localhost")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("16379")
	RedisUser = file.Section("redis").Key("RedisUser").MustString("")
	RedisPassWord = file.Section("redis").Key("RedisPassWord").MustString("")
}

func LoadQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}

func LoadScan(file *ini.File) {
	Account = file.Section("scan").Key("Account").String()
	Pwd = file.Section("scan").Key("Pwd").String()
	Url = file.Section("scan").Key("Url").String()
}
