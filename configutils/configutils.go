package configutils

import (
	"io/ioutil"
	"os"

	"github.com/forallthis/neptune/pathutils"
)

// ReadConfig 读取配置文件
func ReadConfig(configFile string) ([]byte, error) {
	confDir := pathutils.GetExeDir()
	if _, err := os.Stat(confDir + "/" + configFile); os.IsNotExist(err) {
		confDir, _ = os.Getwd() // 执行命令的当前目录
		if _, err := os.Stat(confDir + "/" + configFile); os.IsNotExist(err) {
			panic(err)
		}
	}
	return ioutil.ReadFile(confDir + "/" + configFile)
}

// GetConfigFilePath 获取配置文件目录
func GetConfigFilePath(configFile string) string {
	confDir := pathutils.GetExeDir()
	if _, err := os.Stat(confDir + "/" + configFile); os.IsNotExist(err) {
		confDir, _ = os.Getwd() // 执行命令的当前目录
		if _, err := os.Stat(confDir + "/" + configFile); os.IsNotExist(err) {
			panic(err)
		}
	}
	return confDir + "/" + configFile
}
