package logging

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type logFormat struct {
	TimestampFormat string
}

func init() {
	logrus.SetReportCaller(true)
	formatter := logFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	logrus.SetFormatter(&formatter)

	//ログ出力ファイルの設定
	f, err := os.Create(fmt.Sprintf("%s/src/github.com/taniwhy/mochi-match-rest/config/log/access.%s.log",
		os.Getenv("GOPATH"), time.Now().Format("2006-01-02-15:04:05")))
	if err != nil {
		panic(err.Error())
	}

	gin.DefaultWriter = io.MultiWriter(f)

	f, err = openFile(fmt.Sprintf("%s/src/github.com/taniwhy/mochi-match-rest/config/log/server.%s.log",
		os.Getenv("GOPATH"), time.Now().Format("2006-01-02-15:04:05")))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, f))

	//ログレベルの設定
	logrus.SetLevel(logrus.InfoLevel)

}

//openFile ログを出力するファイルを設定する。
//ファイルが存在する場合、ファイルにログを追記。
//ファイルが存在しない場合、ファイルを作成し、ログを出力。
func openFile(fileName string) (*os.File, error) {
	if exists(fileName) {
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0777)
		return f, err
	}
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	return f, err
}

//formatFilePath ログに記載するファイル名の抽出
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]

}

//exists　ファイルが存在するか確認する。
func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

//Format ログの形式を設定
func (f *logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("]:")
	b.WriteString(entry.Time.Format(f.TimestampFormat))

	b.WriteString(" [")
	b.WriteString(formatFilePath(entry.Caller.File))
	b.WriteString(":")
	fmt.Fprint(b, entry.Caller.Line)
	b.WriteString("] ")

	if entry.Message != "" {
		b.WriteString(" - ")
		b.WriteString(entry.Message)
	}

	if len(entry.Data) > 0 {
		b.WriteString(" || ")
	}
	for key, value := range entry.Data {
		b.WriteString(key)
		b.WriteByte('=')
		b.WriteByte('{')
		fmt.Fprint(b, value)
		b.WriteString("}, ")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

//SetLevelDebug Debugレベルに設定
func SetLevelDebug() {
	logrus.SetLevel(logrus.DebugLevel)
}

//SetLevelInfo Set Infoレベルに設定
func SetLevelInfo() {
	logrus.SetLevel(logrus.InfoLevel)
}

// GormLogger : SQLロガー
type GormLogger struct{}

// Print :
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logrus.WithFields(
			logrus.Fields{
				"module":        "gorm",
				"type":          "sql",
				"rows_returned": v[5],
				"src":           v[1],
				"values":        v[4],
				"duration":      v[2],
			},
		).Info(v[3])
	case "log":
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
