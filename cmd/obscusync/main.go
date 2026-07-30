package main

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/dxo1a/obscusync/internal/app"
	"github.com/inconshreveable/mousetrap"
)

var Version = "dev"

func main() {
	if runtime.GOOS == "windows" && mousetrap.StartedByExplorer() {
		// restart ourselves via cmd /k so that the window does not close
		cmd := exec.Command("cmd", "/k", os.Args[0])
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: false}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		_ = cmd.Run()
		return
	}

	app.Run()
}
