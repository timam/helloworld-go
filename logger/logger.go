package logger

import (
	"flag"
	"log"
	"os"
)

var (
	Log      *log.Logger
)


func init() {
	// set location of log file build.Default.GOPATH + "hello.log"
	var logpath = "/log/helloworld-go.log"

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.Ldate | log.Ltime )
	Log.Println("LogFile : " + logpath)
}