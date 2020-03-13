package gostring

import (
	"strings"
)

type Stringmodel struct {
	Value string
	// 去掉头尾的字符串
	trimcutset string
	first      bool
}

func NewStringmodel(s string) *Stringmodel {
	s2 := &Stringmodel{Value: s}
	s2.trimcutset = " "
	return s2
}

/**
获取字符串的值
*/
func (this *Stringmodel) Getvalue() string {
	return this.Value
}

/**
初始化字符串的值
*/
func (this *Stringmodel) Setvalue(s string) *Stringmodel {
	this.Value = s
	return this
}

func (this *Stringmodel) GetTrimcutset() string {
	return this.trimcutset
}

func (this *Stringmodel) SetTrimcutset(trimcutset string) *Stringmodel {
	this.trimcutset = trimcutset
	return this
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
