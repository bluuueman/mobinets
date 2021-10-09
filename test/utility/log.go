package utility

/*
Log
Check Error
*/
import (
	"log"
	"os"
	"time"
)

func IsErr(err error, msg string) bool {
	if err != nil {
		log.Println("ERROR: "+msg+"\n", err)
		return true
	}
	return false
}

func WriteLog(dir string, buf string) {
	file, oErr := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 666)
	if !IsErr(oErr, "File Open Failed!") {
		curTime := time.Now().Format("2006/01/02 15:04:05") + "\n"
		_, wErr := file.Write([]byte(curTime + buf + "\n"))
		IsErr(wErr, "")
	}
	file.Close()

}
