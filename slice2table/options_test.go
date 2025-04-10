package slice2table_test

import (
	"testing"

	"github.com/go-xlan/go-webpage/slice2table"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewOptions(t *testing.T) {
	options := slice2table.NewOptions()
	t.Log(neatjsons.S(options))
	require.Equal(t, "table", options.TagName)
}
