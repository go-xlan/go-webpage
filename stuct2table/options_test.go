package struct2table

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewOptions(t *testing.T) {
	options := NewOptions()
	t.Log(neatjsons.S(options))
	require.Equal(t, "table", options.TagName)
}
