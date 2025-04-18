package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/firefoxopen"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/yyle88/must"
)

type AdminInfo struct {
	Name string
	From string
}

type GuestInfo struct {
	Name string
	From string
}

func main() {
	admins := []*AdminInfo{
		newAdminInfo(),
		newAdminInfo(),
		newAdminInfo(),
	}
	page1 := utils.NewPage("admins", slice2table.NewTable(admins))

	guests := []*GuestInfo{
		newGuestInfo(),
		newGuestInfo(),
		newGuestInfo(),
	}
	page2 := utils.NewPage("guests", slice2table.NewTable(guests))

	firefoxDraw := firefoxopen.NewFirefoxDraw()
	defer firefoxDraw.Close(time.Minute)
	firefoxDraw.ShowInNewTabs(page1, page2)
	fmt.Println("done")
}

func newAdminInfo() *AdminInfo {
	admin := &AdminInfo{}
	must.Done(gofakeit.Struct(admin))
	return admin
}

func newGuestInfo() *GuestInfo {
	guest := &GuestInfo{}
	must.Done(gofakeit.Struct(guest))
	return guest
}
