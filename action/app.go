package action

import (
	"golang-test/test01/example01/autoGit/slack"
	"os"
	"os/exec"
	"syscall"
)

func GetProperties() string {
	rst := "D:/workspace_go/golang-test"
	return rst
}

func ActionPull() {
	err := os.Chdir(GetProperties())
	if err != nil {
		return
	}
	result, err := execCommandOutput("git", "pull")

	slack.SendMessageToSlack("실행결과 ::\n" + result)

}

func execCommandOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	return string(output), err
}
