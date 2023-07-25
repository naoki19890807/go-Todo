package utils

import (
	"io"
	"log"
	"os"
)

// logファイル生成
func LoggingSettings(logFile string) {
	//読み書き、作成、追記、権限
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logfile err=%s", err.Error())
	}
	//logの書き込み先を標準とlogファイルとする。
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//logのフォーマットを指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
