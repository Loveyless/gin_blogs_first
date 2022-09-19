package controller

import (
	"gin_blogs_first/dao"
	"gin_blogs_first/model"
	"gin_blogs_first/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEditor(c *gin.Context) {
	editorAllList := dao.Mgr.GetAllEditor()
	c.JSON(200, gin.H{
		"status":  200,
		"message": "查询成功",
		"data":    editorAllList,
	})
}

func GetEditor(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	e := dao.Mgr.GetEditor(pid)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "查询成功",
		"data":    e,
	})
}

func AddEditor(c *gin.Context) {
	editor := model.Editor{}
	//校验
	if err := c.ShouldBind(&editor); err != nil {
		c.JSON(200, gin.H{
			"status":  400,
			"message": validator.Translate(err),
			"data":    "",
		})
		return
	}

	//逻辑
	b := dao.Mgr.AddEditor(&editor)
	if b {
		c.JSON(200, gin.H{
			"message": "添加成功",
			"status":  200,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "添加失败",
			"status":  400,
		})
	}
}
