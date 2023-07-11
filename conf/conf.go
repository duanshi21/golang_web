package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig 读取配置文件
func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	// 判断配置文件是否读取错误
	if err != nil {
		panic(fmt.Sprintf("Load Config Error：%s", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}
