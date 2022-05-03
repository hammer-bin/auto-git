package action

import (
	"auto-git/slack"
	"os"
	"os/exec"
	"syscall"
)

func GetProperties() string {
	rst := "D:/workspace_go/auto-git"
	return rst
}

func ProcessPull() {
	err := os.Chdir(GetProperties())
	if err != nil {
		return
	}
	result, err := execCommandOutput("git", "pull")

	err = slack.SendMessageToSlack("실행결과 ::\n" + result)
	if err != nil {
		return
	}

}

func execCommandOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}
