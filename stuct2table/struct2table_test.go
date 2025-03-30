package struct2table_test

import (
	"testing"

	struct2table "github.com/go-xlan/go-web-table/stuct2table"
	"github.com/stretchr/testify/require"
)

func TestNewTable(t *testing.T) {
	type UserType struct {
		Name string `table:"姓名"`
		Age  int    `table:"年龄"`
	}

	var users = []*UserType{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}
	webTable := struct2table.NewTable(users, struct2table.NewOptions())
	t.Log(webTable)

	require.Equal(t, `<table border="1"><tr><th>姓名</th><th>年龄</th></tr><tr><td>Alice</td><td>25</td></tr><tr><td>Bob</td><td>30</td></tr></table>`, webTable)
}

func TestNewTable0(t *testing.T) {
	type StudentType struct {
		Name string `table:"姓名"`
		Rank int    `table:"排名"`
	}

	var students []*StudentType
	webTable := struct2table.NewTable(students, struct2table.NewOptions())
	t.Log(webTable)

	require.Equal(t, `<table border="1"><tr><th>姓名</th><th>排名</th></tr></table>`, webTable)
}

func TestNewTable1(t *testing.T) {
	type StudentType struct {
		Name string `TH:"姓名"`
		Rank int    `TH:"排名"`
	}

	var students = []*StudentType{{Name: "LuoLuo", Rank: 1}}
	webTable := struct2table.NewTable(students, struct2table.NewOptions().WithTagName("TH"))
	t.Log(webTable)

	require.Equal(t, `<table border="1"><tr><th>姓名</th><th>排名</th></tr><tr><td>LuoLuo</td><td>1</td></tr></table>`, webTable)
}
