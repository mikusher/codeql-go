// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import "html"

func TaintStepTest_HtmlEscapeString_B0I0O0(sourceCQL interface{}) interface{} {
	fromString656 := sourceCQL.(string)
	intoString414 := html.EscapeString(fromString656)
	return intoString414
}

func TaintStepTest_HtmlUnescapeString_B0I0O0(sourceCQL interface{}) interface{} {
	fromString518 := sourceCQL.(string)
	intoString650 := html.UnescapeString(fromString518)
	return intoString650
}

func RunAllTaints_Html() {
	{
		source := newSource(0)
		out := TaintStepTest_HtmlEscapeString_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_HtmlUnescapeString_B0I0O0(source)
		sink(1, out)
	}
}