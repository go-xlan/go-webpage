package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

type User struct {
	ID        string    // 用户 ID
	Username  string    // 用户名
	Email     string    // 邮箱
	Role      string    // 角色（如 admin, guest, student）
	IsActive  bool      // 是否激活
	CreatedAt time.Time // 创建时间
}

type Order struct {
	ID        string    // 订单 ID
	UserID    string    // 用户 ID
	Total     float64   // 总金额
	Status    string    // 状态（如 pending, shipped）
	OrderedAt time.Time // 下单时间
}

func main() {
	users := []*User{
		newUser(),
		newUser(),
		newUser(),
	}
	page1 := utils.NewPage("users", slice2table.NewTable(users))

	orders := []*Order{
		newOrder(),
		newOrder(),
		newOrder(),
	}
	page2 := utils.NewPage("orders", slice2table.NewTable(orders))

	commandConfig := osexec.NewOsCommand().WithDebug()
	w3mopenpage.Show(commandConfig, page1)
	w3mopenpage.Show(commandConfig, page2)
	fmt.Println("done")
}

func newUser() *User {
	user := &User{}
	must.Done(gofakeit.Struct(user))
	return user
}

func newOrder() *Order {
	order := &Order{}
	must.Done(gofakeit.Struct(order))
	return order
}
