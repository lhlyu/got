package gen

import (
	"github.com/lhlyu/got/v2/db/core"
	"github.com/lhlyu/yutil/v2"
)

const g = `
// 获取
func (this *{{.}}) Get() error{
    if this.Id == 0{
        return MissPkErr
    }
    return common.DB.First(this,this.Id).Error
}

// 分页查询
func (this *{{.}}) Query(rs interface{},page *Page,whr map[string]interface{}, order string) error{
    var total int
    if err := common.DB.Model(this).Where(whr).Count(&total).Error;err != nil{
        return err
    }
    page.SetTotal(total)
    return common.DB.Where(whr).Offset(page.StartRow).Limit(page.PageSize).Order(order).Find(rs).Error
}

// 添加
func (this *{{.}}) Add() error{
    return common.DB.Create(this).Error
}

// 删除
func (this *{{.}}) Del() error{
    if this.Id == 0{
        return MissPkErr
    }
    return common.DB.Unscoped().Delete(this).Error
}

// 更新
func (this *{{.}}) Update(whr map[string]interface{}) error{
    if this.Id == 0{
        return MissPkErr
    }
    if whr == nil{
        return common.DB.Model(this).Updates(this).Error
    }
    return common.DB.Model(this).Updates(whr).Error
}

`

var GORM_CURD_FUNC = func(tab *core.Table) string {
	return yutil.String.TemplateParse(g, yutil.String.BigCamelCase(tab.Name))
}
