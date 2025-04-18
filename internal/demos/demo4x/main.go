package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/must"
)

type Product struct {
	ID          string    // 商品 ID
	Name        string    // 商品名称
	Description string    // 描述
	Price       float64   // 价格
	Stock       int       // 库存量
	Category    string    // 分类
	CreatedAt   time.Time // 创建时间
}

type Customer struct {
	ID        string    // 客户 ID
	Name      string    // 姓名
	Email     string    // 邮箱
	Phone     string    // 电话
	Address   string    // 地址
	CreatedAt time.Time // 创建时间
}

func main() {
	products := []*Product{
		newProduct(),
		newProduct(),
		newProduct(),
	}
	page1 := utils.NewPage("products", slice2table.NewTable(products))

	customers := []*Customer{
		newCustomer(),
		newCustomer(),
		newCustomer(),
	}
	page2 := utils.NewPage("customers", slice2table.NewTable(customers))

	w3mDrawPage := w3mopenpage.NewW3mDrawPage()
	defer w3mDrawPage.Close()
	w3mDrawPage.ShowPages(page1, page2)
	fmt.Println("done")
}

func newProduct() *Product {
	product := &Product{}
	must.Done(gofakeit.Struct(product))
	return product
}

func newCustomer() *Customer {
	customer := &Customer{}
	must.Done(gofakeit.Struct(customer))
	return customer
}
