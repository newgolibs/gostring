package gostring

type Stringmodel struct {
	Value string
	// 去掉头尾的字符串
	trimcutset string
}

func (this *Stringmodel) GetTrimcutset() string {
	return this.trimcutset
}

func (this *Stringmodel) SetTrimcutset(trimcutset string) *Stringmodel {
	this.trimcutset = trimcutset
	return this
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