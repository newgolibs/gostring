package gostring

type Stringmodel_interface interface {
	//在原来的基础上追加字符串
	Append(s string) *Stringmodel
	Invoke() string
}
