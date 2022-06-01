package onebot

import (
	"bot_app/bot_myqq/core"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func INFO(s string, v ...interface{}) {
	core.OutPutLog("[INFO] " + fmt.Sprintf(s, v...))
}

func WARN(s string, v ...interface{}) {
	core.OutPutLog("[WARN] " + fmt.Sprintf(s, v...))
}

func DEBUG(s string, v ...interface{}) {
	core.OutPutLog("[DEBUG] " + fmt.Sprintf(s, v...))
}

func ERROR(s string, v ...interface{}) {
	core.OutPutLog("[ERROR] " + fmt.Sprintf(s, v...))
}

func META(s string, v ...interface{}) {
	core.OutPutLog("[META] " + fmt.Sprintf(s, v...))
}

func TEST(s string, v ...interface{}) {
	core.OutPutLog("[TEST] " + fmt.Sprintf(s, v...))
}

func PathExecute() string {
	dir, err := os.Getwd()
	if err != nil {
		ERROR("判断当前运行路径失败")
	}
	return dir + "/"
}

func CreatePath(path string) {
	if !PathExists(path) {
		length := len(path)
		switch {
		case path[length-1:] != "/":
			path = path[:strings.LastIndex(path, "/")]
		case path[length-1:] != "\\":
			path = path[:strings.LastIndex(path, "\\")]
		default:
			//
		}
		err := os.MkdirAll(path, 0644)
		if err != nil {
			ERROR("生成应用目录失败")
		}
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ProtectRun(entry func(), label string) {
	defer func() {
		err := recover()
		if err != nil {
			ERROR("[协程] %v协程发生了不可预知的错误，请在GitHub提交issue：%v", label, err)
			buf := make([]byte, 1<<16)
			runtime.Stack(buf, true)
			ERROR("traceback:\n%v", string(buf))
		}
	}()
	entry()
}
