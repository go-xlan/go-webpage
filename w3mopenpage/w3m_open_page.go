package w3mopenpage

import (
	"os/exec"
	"strings"

	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func Open(command *osexec.OsCommand, link string) {
	output := rese.V1(command.Exec("w3m", "-dump", link))
	zaplog.SUG.Debug("[page]:", "\n", string(output), "\n", "----")
}

func Show(command *osexec.OsCommand, page string) {
	output := rese.V1(command.ExecWith("w3m", []string{"-T", "text/html", "-dump"}, func(command *exec.Cmd) {
		command.Stdin = strings.NewReader(page)
	}))
	zaplog.SUG.Debug("[page]:", "\n", string(output), "\n", "----")
}
