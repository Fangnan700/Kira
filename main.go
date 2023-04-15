package main

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"io"
	"kira/model"
	"kira/utils"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	logFile    *os.File
	logWriter  io.Writer
	err        error
	configFile []byte
	config     model.Config
)

func init() {

	// 配置log
	logPath, _ := filepath.Abs("kira-log/kira-app.log")
	logFile, _ = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logWriter = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)

	// 读取配置文件
	configPath, _ := filepath.Abs("kira-config/config.yml")
	configFile, err = os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if yaml.Unmarshal(configFile, &config) != nil {
		log.Fatalln(err)
	}
}

func main() {
	// 读取用户输入
	for {
		fmt.Printf("[%s] # ", time.Now().Format("15:04:05"))
		var sendContent string
		_, _ = fmt.Scanln(&sendContent)
		if sendContent == "拜拜" {
			color.Set(color.FgHiCyan)
			fmt.Printf("[%s] $ 喵～主人再见♥", time.Now().Format("15:04:05"))
			time.Sleep(1000 * time.Millisecond)
			color.Unset()
			utils.ClearConsole()
			return
		}
		utils.SendMessage(sendContent)
	}
}
