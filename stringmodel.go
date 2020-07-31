package gostring

import (
	"strings"
)
/**
构造函数
 */
func NewStringmodel(s string) *Stringmodel {
	Stringmodel_obj := new(Stringmodel)
	Stringmodel_obj.Value = s
	// 默认字符串trim的字符内容是空格
	Stringmodel_obj.trimcutset = " "
	return Stringmodel_obj
}

/**
追加字符串
*/
func (this *Stringmodel) Append(s string) *Stringmodel {
	this.Value = this.Value + s
	return this
}

/**
首字母大写
*/
func (this *Stringmodel) Ucfirst() *Stringmodel {
	s := []byte(this.Value)
	this.Value = strings.Replace(this.Value, string(s[0]), strings.ToUpper(string(s[0])), 1)
	return this
}

/**
全部转换成大写
*/
func (this *Stringmodel) Toupcase() *Stringmodel {
	this.Value = strings.ToUpper(this.Value)
	return this
}

/**
全部转换成小写
*/
func (this *Stringmodel) Tolowercase() *Stringmodel {
	this.Value = strings.ToLower(this.Value)
	return this
}

/**
去掉头尾指定的字符串,默认去掉空格
*/
func (this *Stringmodel) Trim() *Stringmodel {
	this.Value = strings.Trim(this.Value, this.trimcutset)
	return this
}

/**
所有的计算合并一起执行
*/
func (this *Stringmodel) Invoke() string {
	return this.Value
}
