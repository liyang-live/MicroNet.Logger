// LogModel
package models

type LogModel struct {
	//编号
	Id string

	//系统
	System string

	//模块
	Module string

	//方法名
	FuncName string

	//参数
	FuncParameter string

	//记录时间
	LogTime string

	//日志内容
	LogContent string
}
