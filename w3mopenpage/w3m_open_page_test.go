package w3mopenpage_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

func TestShow(t *testing.T) {
	w3mopenpage.Show(osexec.NewOsCommand().WithDebug(), utils.NewPage("标题1", "内容1"))
}

type Account struct {
	Username string
	Password string
	Nickname string
	IsVip    bool
}

func newAccount() *Account {
	account := &Account{}
	must.Done(gofakeit.Struct(account))
	return account
}

func TestShowPage(t *testing.T) {
	w3mopenpage.Show(osexec.NewOsCommand().WithDebug(), utils.NewPage("accounts", slice2table.NewTable([]*Account{
		newAccount(),
		newAccount(),
		newAccount(),
	})))
}
