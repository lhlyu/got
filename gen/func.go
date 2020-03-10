package gen

import (
	"fmt"
	"github.com/lhlyu/got/db/core"
	"github.com/lhlyu/yutil/v2"
)

const g = `
func (%s) Name() string {
	return "%s"
}
`

var GORM_CURD = func(tab *core.Table) string {
	return fmt.Sprintf(g, yutil.String.BigCamelCase(tab.Name), tab.Name)
}
