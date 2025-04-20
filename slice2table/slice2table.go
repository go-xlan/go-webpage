package slice2table

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/yyle88/must"
	"github.com/yyle88/printgo"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// NewTable 将结构体指针切片转换为 HTML 表格，支持自定义标签名
func NewTable[T any](objects []*T) string {
	return GenTable(objects, NewOptions())
}

func GenTable[T any](objects []*T, options *Options) string {
	columnNames, dataRows := extract(objects, must.Nice(options))

	tableHeadLine := newHeadLine(columnNames)
	tableBodyRows := newDataRows(dataRows)

	return `<table border="1">` + tableHeadLine + strings.Join(tableBodyRows, "") + `</table>`
}

func extract[T any](objects []*T, options *Options) ([]string, [][]string) {
	structType := syntaxgo_reflect.GetTypeV2[T]()
	must.Same(structType.Kind(), reflect.Struct)

	columnNames := extractKeys(structType, options)
	dataRows := extractRows(structType, objects, options)

	return columnNames, dataRows
}

func extractKeys(elemType reflect.Type, options *Options) []string {
	var columnNames []string
	for idx := 0; idx < elemType.NumField(); idx++ {
		field := elemType.Field(idx)
		if field.PkgPath == "" { // 只处理导出字段
			tagValue := field.Tag.Get(options.TagName)
			if tagValue == "" {
				tagValue = field.Name
			}
			columnNames = append(columnNames, tagValue)
		}
	}
	return columnNames
}

func extractRows[T any](elemType reflect.Type, objects []*T, options *Options) [][]string {
	var dataRows [][]string
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
		dataRows = append(dataRows, row)
	}
	return dataRows
}

func newHeadLine(keys []string) string {
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

func newDataRows(data [][]string) []string {
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
