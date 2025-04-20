package struct2table

import (
	"github.com/go-xlan/go-webpage/slice2table"
)

func NewTable[T any](object *T) string {
	return GenTable(object, slice2table.NewOptions())
}

func GenTable[T any](object *T, options *slice2table.Options) string {
	return slice2table.GenTable([]*T{object}, options)
}
