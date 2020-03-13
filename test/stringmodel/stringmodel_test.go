package stringmodel

import (
	"fmt"
	"github.com/newgolibs/gostring"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	str_obj *gostring.Stringmodel
)

/**
追加字符串
*/
func TestAppend(t *testing.T) {
	old_str := "xxx"
	str_obj = gostring.NewStringmodel(old_str)
	fmt.Printf("%+v\n", str_obj)
	// 断言 : 字符串赋值成功
	assert.Equal(t, old_str, str_obj.Getvalue())
	old_str_append := " add new"
	str_obj.Append(old_str_append)
	fmt.Printf("%+v\n", str_obj)
	// 断言:字符串被追加成功
	assert.Equal(t, old_str+old_str_append, str_obj.Invoke())
}

/**
全部转换成大写
*/
func TestToupcase(t *testing.T) {
	old_str := "xxx"
	str_obj = gostring.NewStringmodel(old_str)
	str_obj.Toupcase()
	fmt.Printf("%+v\n", str_obj)
	// 断言:字符串全部变成大写了
	assert.Equal(t, "XXX", str_obj.Invoke())
}

/**
全部转换成小写
*/
func TestTolowercase(t *testing.T) {
	old_str := "AbC"
	str_obj = gostring.NewStringmodel(old_str)
	str_obj.Tolowercase()
	fmt.Printf("%+v\n", str_obj)
	// 断言:字符串全部变成小写了
	assert.Equal(t, "abc", str_obj.Invoke())
}

/**
去掉头尾指定的字符串,默认去掉空格
*/
func TestTrim_默认指定去掉空格(t *testing.T) {
	old_str := "   AbC    "
	str_obj = gostring.NewStringmodel(old_str)
	fmt.Printf("%+v\n", str_obj)
	str_obj.Trim()
	fmt.Printf("%+v\n", str_obj)
	// 断言:头尾的空格全部被除掉了
	assert.Equal(t, "AbC", str_obj.Invoke())
}

/**
去掉头尾指定的字符串,默认去掉空格
*/
func TestTrim_用户自己指定要去除的内容(t *testing.T) {
	str_obj2 := gostring.NewStringmodel("AbbbbbA")
	fmt.Printf("%+v\n", str_obj2)
	str_obj2.
		SetTrimcutset("A").
		Trim()
	fmt.Printf("%+v\n", str_obj2)
	// 断言:头尾的空格全部被除掉了
	assert.Equal(t, "bbbbb", str_obj2.Invoke())
}

/**
首字母大写
*/
func TestUcfirst(t *testing.T)  {
	old_str := "hello how are you"
	str_obj = gostring.NewStringmodel(old_str)
	str_obj.Ucfirst()
	fmt.Printf("%+v\n", str_obj)
	assert.Equal(t,"Hello how are you",str_obj.Invoke())
}