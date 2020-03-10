package core

import "strings"

type Table struct {
	Name        string
	Comment     string
	ColumnsMap  map[string]*Column
	Columns     []*Column
	PrimaryKeys []string
}

func NewTable() *Table {
	return &Table{
		ColumnsMap: make(map[string]*Column),
	}
}

func (t *Table) AddColumn(c *Column) *Table {
	t.Columns = append(t.Columns, c)
	t.ColumnsMap[c.Name] = c
	return t
}

func (t *Table) AddColumns(cc []*Column) *Table {
	t.Columns = append(t.Columns, cc...)
	//for _, v := range cc {
	//	t.ColumnsMap[v.Name] = v
	//}
	return t
}

func (t *Table) AddPk(pk string) *Table {
	t.PrimaryKeys = append(t.PrimaryKeys, pk)
	return t
}

func (t *Table) AddComment(comment string) *Table {
	comment = strings.TrimSpace(comment)
	t.Comment = comment
	return t
}
