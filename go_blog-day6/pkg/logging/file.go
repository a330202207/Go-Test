package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

//获取日志文件
func getLogFilePath() string {
	dir, _ := os.Getwd()
	path := dir + "/" + LogSavePath
	log.Printf("Log file path : %s", path)
	return fmt.Sprintf("%s", path)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

//打开日志文件
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err): //是否存在
		mkDir()
	case os.IsPermission(err): //是否有权限
		log.Fatalf("Permission :%v", err)
	}

	// 在写入时将数据追加到文件中|如果不存在，则创建一个新文件|以只写模式打开文件
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

//创建目录
func mkDir() {
	// 返回与当前目录对应的根路径名
	dir, _ := os.Getwd()

	//创建对应的目录以及所需的子目录
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
