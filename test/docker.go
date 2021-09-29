package main

//for test
import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func phase1(output *bytes.Buffer) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", "docker images")
	cmd.Stdout = output
	log.Println("Docker Cmd")
	cmdRunErr := cmd.Start()
	if cmdRunErr != nil {
		log.Println(cmdRunErr)
		return nil
	}
	return cmd
}

func writeLog(dir string, buf []byte) {
	file, _ := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 666)
	file.Write(buf)
	file.Close()
}
func main() {
	var output bytes.Buffer
	cmd := phase1(&output)
	if cmd != nil {
		cmd.Wait()
		writeLog("docker.txt", []byte(output.String()))

	}

}
