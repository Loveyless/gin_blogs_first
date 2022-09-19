package dao

import (
	"gin_blogs_first/model"
)

//创建博客文章
func (mgr *manager) AddEditor(editor *model.Editor) bool {
	d := mgr.db.Create(editor)
	if d.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

//查询所有博客文章
func (mgr *manager) GetAllEditor() []model.Editor {
	var editorList []model.Editor
	mgr.db.Find(&editorList)
	return editorList
}

//通过pid查询博客文章
func (mgr *manager) GetEditor(pid int) model.Editor {
	var editor model.Editor
	mgr.db.Where("id = ?", pid).First(&editor)
	return editor
}
