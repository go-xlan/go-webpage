package w3mopenpage_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/must"
)

var caseDrawPage *w3mopenpage.W3mShowPage

func TestMain(m *testing.M) {
	w3mDrawPage := w3mopenpage.NewW3mDrawPage()
	defer w3mDrawPage.Close()

	caseDrawPage = w3mDrawPage
	m.Run()
}

func TestW3mShowPage_Show(t *testing.T) {
	caseDrawPage.Show(utils.NewPage("标题1", "内容1"))
}

type Employee struct {
	ID       string
	Nickname string
	Role     string
}

func newEmployee() *Employee {
	employee := &Employee{}
	must.Done(gofakeit.Struct(employee))
	return employee
}

func TestW3mShowPage_ShowPage(t *testing.T) {
	caseDrawPage.Show(utils.NewPage("employees", slice2table.NewTable([]*Employee{
		newEmployee(),
		newEmployee(),
		newEmployee(),
	})))
}
