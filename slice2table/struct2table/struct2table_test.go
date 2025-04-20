package struct2table_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-xlan/go-webpage/slice2table/struct2table"
	"github.com/stretchr/testify/require"
)

func TestNewTable(t *testing.T) {
	type DemoType struct {
		Name string
		Code string
		Rank int
	}

	demo := &DemoType{}
	require.NoError(t, gofakeit.Struct(demo))

	res := struct2table.NewTable(demo)
	t.Log(res)
}
