package pathutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/forallthis/neptune/stringutils"
)

// GetParentDirectory 获取父路径
func GetParentDirectory(dirctory string) string {
	return stringutils.Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// GetExeDir 可执行文件的目录
func GetExeDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
