package main

import (
	"surge/cmd"
	"surge/cmdutil/verbose"

	"github.com/spf13/viper"
)

func main() {

	cmd.Execute()

	//load()
}

func load() {

	// 设置文件名
	viper.SetConfigName(".surge.conf") // name of config file (without extension)
	// 设置配置文件的类型
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	// 设置配置文件路径为当前工作目录
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/") // 设置配置文件的搜索目录
	if err := viper.ReadInConfig(); err != nil {
		verbose.Println(err)

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// verbose.Println("配置文件未找到，载入默认配置")
		} else {
			// verbose.Println("配置文件已找到，但出现了其他的错误")
		}
	} else {
		// verbose.Println("找到配置文件")
		//server := getConfigKey("server","127.0.0.1")
		//port := getConfigKey("port","6170")
		//key := getConfigKey("x-key","")
		// verbose.Println("key is " + server + ":"+ port+ "@" +key)
	}
}
