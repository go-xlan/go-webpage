package w3mopenpage

import (
	"github.com/go-xlan/go-webpage/gintestpage"
	"github.com/google/uuid"
	"github.com/yyle88/osexec"
)

type W3mShowPage struct {
	service *gintestpage.Service
	command *osexec.OsCommand
}

func NewW3mShowPage(service *gintestpage.Service, command *osexec.OsCommand) *W3mShowPage {
	return &W3mShowPage{
		service: service,
		command: command,
	}
}

func NewW3mDrawPage() *W3mShowPage {
	return NewW3mShowPage(gintestpage.NewService(), osexec.NewOsCommand().WithDebug())
}

func (op *W3mShowPage) Close() {
	op.service.Close()
}

func (op *W3mShowPage) Show(page string) {
	path := uuid.New().String()
	link := op.service.SetPage(path, []byte(page))
	Open(op.command, link)
}

func (op *W3mShowPage) ShowPages(pages ...string) {
	for _, page := range pages {
		path := uuid.New().String()
		link := op.service.SetPage(path, []byte(page))
		Open(op.command, link)
	}
}
