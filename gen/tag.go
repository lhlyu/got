package gen

import (
	"fmt"
	"github.com/lhlyu/got/db/core"
	"github.com/lhlyu/yutil/v2"
)

var JSON_TAG = func(col *core.Column) string {
	return fmt.Sprintf("`json:\"%s\"`", yutil.String.LittleCamelCase(col.Name))
}

var DB_TAG = func(col *core.Column) string {
	return fmt.Sprintf("`db:\"%s\"`", col.Name)
}

var JSON_DB_TAG = func(col *core.Column) string {
	return fmt.Sprintf("`json:\"%s\" db:\"%s\"`", yutil.String.LittleCamelCase(col.Name), col.Name)
}
