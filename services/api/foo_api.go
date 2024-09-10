// Package api -----------------------------
// @author   : deng zhi
// @time     : 2024/9/10 14:47
//
// -------------------------------------------
package api

type Args struct{ Num1, Num2 int }

type FooAPI interface {
	Sum(args Args, reply *int) error
}
