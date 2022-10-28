package jsonize

import (
	"encoding/json"

	"github.com/bytedance/sonic"
)

// Jsonize 返回对象的json内容，无视格式化错误.
func JsonToString(object interface{}, indent bool) string {
	if indent {
		b, _ := json.MarshalIndent(object, "", "  ")
		return string(b)
	}
	b, _ := sonic.Marshal(object)
	return string(b)
}
