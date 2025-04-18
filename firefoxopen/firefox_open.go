package firefoxopen

import (
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func OpenInNewWindows(command *osexec.OsCommand, urls ...string) {
	Open(command, urls, "--new-window") //打开若干个新窗口以打开若干个网页
}

func OpenInNewTabs(command *osexec.OsCommand, urls ...string) {
	Open(command, urls, "--new-tab") //打开若干个新标签以打开若干个网页
}

func Open(command *osexec.OsCommand, urls []string, openOption string) {
	args := []string{must.Nice(openOption)}
	args = append(args, urls...)
	output := rese.V1(command.Exec("firefox", args...))
	zaplog.SUG.Debugln(string(output))
}
