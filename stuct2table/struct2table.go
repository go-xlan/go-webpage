package struct2table

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/printgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// NewTable 将结构体指针切片转换为 HTML 表格，支持自定义标签名
func NewTable[T any](objects []*T, options *Options) string {
	keys, rows := extract(objects, options)

	tableTrTHs := newTroTHs(keys)
	tableTrTDs := newTrnTDs(rows)

	return `<table border="1">` + tableTrTHs + strings.Join(tableTrTDs, "") + `</table>`
}

func extract[T any](objects []*T, options *Options) ([]string, [][]string) {
	elemType := syntaxgo_reflect.GetTypeV2[T]()
	must.Same(elemType.Kind(), reflect.Struct)

	keys := extractKeys(elemType, options)
	data := extractRows(elemType, objects, options)

	return keys, data
}

func extractKeys(elemType reflect.Type, options *Options) []string {
	var keys []string
	for idx := 0; idx < elemType.NumField(); idx++ {
		field := elemType.Field(idx)
		if field.PkgPath == "" { // 只处理导出字段
			tagValue := field.Tag.Get(options.TagName)
			if tagValue == "" {
				tagValue = field.Name
			}
			keys = append(keys, tagValue)
		}
	}
	return keys
}

func extractRows[T any](elemType reflect.Type, objects []*T, options *Options) [][]string {
	var rows [][]string
	for _, item := range objects {
		elem := reflect.ValueOf(item).Elem()
		var row []string
		for idx := 0; idx < elemType.NumField(); idx++ {
			field := elemType.Field(idx)
			if field.PkgPath == "" { // 只处理导出字段
				value := elem.Field(idx).Interface()
				row = append(row, formatValue(value, options))
			}
		}
		rows = append(rows, row)
	}
	return rows
}

func newTroTHs(keys []string) string {
	sb := printgo.NewPTS()
	if len(keys) > 0 {
		sb.WriteString("<tr>")
		for _, k := range keys {
			sb.Fprintf("<th>%s</th>", k)
		}
		sb.WriteString("</tr>")
	}
	return sb.String()
}

func newTrnTDs(data [][]string) []string {
	var rows []string
	for _, row := range data {
		sb := printgo.NewPTS()
		sb.WriteString("<tr>")
		for _, v := range row {
			sb.Fprintf("<td>%s</td>", v)
		}
		sb.WriteString("</tr>")
		rows = append(rows, sb.String())
	}
	return rows
}

func formatValue(value interface{}, options *Options) string {
	switch v := value.(type) {
	case string:
		return v
	case float64, float32:
		return fmt.Sprintf("%.2f", v)
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%d", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
