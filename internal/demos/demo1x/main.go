package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/firefoxopen"
	"github.com/go-xlan/go-webpage/gintestpage"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/google/uuid"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

type StudentScore struct {
	Name  string
	Age   int
	Score float64
	Rank  uint32
}

func main() {
	service := gintestpage.NewService()
	defer service.Close()

	scores := []*StudentScore{
		newStudentScore(),
		newStudentScore(),
		newStudentScore(),
	}
	page := utils.NewPage("scores", slice2table.NewTable(scores))
	path := uuid.New().String()
	link := service.SetPage(path, []byte(page))
	fmt.Println("link:", link)

	firefoxopen.OpenInNewTabs(osexec.NewOsCommand().WithDebug(), link)
	fmt.Println("done")
	time.Sleep(time.Second * 3)
}

func newStudentScore() *StudentScore {
	score := &StudentScore{}
	must.Done(gofakeit.Struct(score))
	return score
}
