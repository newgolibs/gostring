package Stringmodel

import "github.com/newgolibs/gostring"

type stringmodel_interface interface {
	//在原来的基础上追加字符串
	Append(s string) *gostring.Stringmodel
	Ucfirst() *gostring.Stringmodel
	Invoke() string
}
