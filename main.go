package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	syscall.Syscall(syscall.
		NewLazyDLL("kernel32.dll").
		NewProc("AttachConsole").
		Addr(),
		1, uintptr(^uint32(0)), 0, 0)

	inst := flag.Bool("inst", false, "install jre mode")
	flag.Parse()

	if *inst {
		RequestJRE()
		return
	}

	if len(config.String("Core.jar")) == 0 {
		config.Set("Core.jar", " ")
	}

	if config.String("Test.update") == "true" {
		CallRequestJRE()
		RunJar(`.\jre`)
		return
	}

	java := config.String("Java.home")
	if len(java) > 0 {
		RunJar(java)
		return
	}

	java = SearchLocalJava()
	if len(java) > 0 {
		RunJar(java)
		return
	}

	CallRequestJRE()
	RunJar(`.\jre`)
}

func RunJar(java string) {
	config.Set("Java.home", java)
	exePath, _ := os.Executable()
	vmOptBuf, err := ioutil.ReadFile(exePath + ".vmoptions")
	vmOpt := string(vmOptBuf)
	if err == nil {
		vmOpt = strings.Replace(vmOpt, "\r\n", " ", -1)
		vmOpt = strings.Replace(vmOpt, "\n", " ", -1)
		vmOpt = strings.Replace(vmOpt, "\r", " ", -1)
	}
	jCmd := exec.Command(java + `\bin\java.exe`, "-jar", config.String("Core.jar"), vmOpt)
	jCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	jCmd.Run()
}

func SearchLocalJava() string {
	javaHome, e := os.LookupEnv("JAVA_HOME")
	if !e {
		return ""
	}
	java := javaHome + `\bin\java.exe`
	if _, err := os.Stat(java); os.IsNotExist(err) {
		return ""
	}
	vCmd := exec.Command(java, "-version")
	vCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	vb, err := vCmd.CombinedOutput()
	if err != nil {
		return ""
	}
	vs := string(vb)
	rxVersion := regexp.MustCompile(`"(\d+\.\d+)`)
	rs := rxVersion.FindAllStringSubmatch(vs, -1)
	if len(rs) == 0 || len(rs[0]) == 1 {
		return ""
	}
	v, err := strconv.ParseFloat(rs[0][1], 64)
	if err != nil {
		return ""
	}
	if v < 1.8 {
		return ""
	}
	return javaHome
}

func CallRequestJRE() {
	exePath, _ := os.Executable()
	iCmd := exec.Command("cmd", "/C", exePath, "-inst")
	iCmd.Dir = filepath.Dir(exePath)
	iCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: false}
	iCmd.Run()
}

func RequestJRE() {
	jre := config.String("Java.source")
	if len(jre) == 0 {
		jre = "http://vec-public-jre.b0.upaiyun.com/jre-{arch}"
		config.Set("Java.source", jre)
	}
	jre = strings.Replace(jre, "{arch}", runtime.GOARCH, -1)
	println("Start downloading dependencies / 开始下载依赖组件 / コンポーネントのダウンロードを開始する")
	Download(jre, "./jre_update")
	println("Unzipping... / 解压中... / 減圧する...")
	Unzip("./jre_update", "./jre")
	os.Remove("./jre_update")
	config.Set("Java.home", `.\jre`)
}