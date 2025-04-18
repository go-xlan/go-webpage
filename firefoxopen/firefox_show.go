package firefoxopen

import (
	"sync"
	"time"

	"github.com/go-xlan/go-webpage/gintestpage"
	"github.com/google/uuid"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

type FirefoxShow struct {
	service  *gintestpage.Service
	command  *osexec.OsCommand
	showTime time.Time
	doneOnce *sync.Once
}

func NewFirefoxShow(service *gintestpage.Service, command *osexec.OsCommand) *FirefoxShow {
	return &FirefoxShow{
		service:  service,
		command:  command,
		showTime: time.Unix(0, 0),
		doneOnce: &sync.Once{},
	}
}

func NewFirefoxDraw() *FirefoxShow {
	return NewFirefoxShow(gintestpage.NewService(), osexec.NewOsCommand().WithDebug())
}

func (op *FirefoxShow) Close(waitTime time.Duration) {
	op.doneOnce.Do(func() {
		time.Sleep(waitTime - time.Since(op.showTime))
		op.service.Close()
	})
}

func (op *FirefoxShow) ShowInNewWindows(pages ...string) {
	op.Show(pages, "--new-window") //打开若干个新窗口以打开若干个网页
}

func (op *FirefoxShow) ShowInNewTabs(pages ...string) {
	op.Show(pages, "--new-tab") //打开若干个新标签以打开若干个网页
}

func (op *FirefoxShow) Show(pages []string, openOption string) {
	must.Have(pages)

	var urls = make([]string, 0, len(pages))
	for _, page := range pages {
		path := uuid.New().String()
		link := op.service.SetPage(path, []byte(page))
		urls = append(urls, link)
	}

	Open(op.command, urls, openOption)

	op.showTime = time.Now()
}
