package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logs = logrus.New()

func Info(info string)  {
	WriteInfo(info)
}

func WriteInfo(info string)  {
	file, err := getFile()
	if err == nil {
		logs.Out = file
	} else {
		logs.Info("Failed to log to file, using default stderr")
	}

	defer file.Close()

	logs.Info(info)
}

func Error(info string)  {
	WriteError(info)
}

func WriteError(info string)  {
	file,err := getFile()
	if err == nil {
		logs.Out = file
	} else {
		logs.Info("Failed to log to file, using default stderr")
	}
	defer file.Close()

	logs.Error(info)
}

func getFile() (*os.File,error) {
	dir := GetLogConfig()
	logDirPath := dir

	if !isExistDir(logDirPath){
		mkdirErr := os.Mkdir(dir,0777)
		if mkdirErr != nil {
			fmt.Printf("创建日志目录失败：%v\n",mkdirErr.Error())
		}
	}

	fileName := "callback_log_"+time.Now().Format("2006-01-02")
	file, err := os.OpenFile(logDirPath+"/"+fileName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	return file,err
}

/**
 * 文件目录是否不存在
 */
func isExistDir(path string) bool {
	_,err := os.Stat(path)
	if err != nil {
		if os.IsExist(err){  // 根据错误类型进行判断
			return true
		}
		return false
	}
	return true
}