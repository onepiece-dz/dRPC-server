// Package services -----------------------------
// @author   : deng zhi
// @time     : 2024/9/10 14:35
//
// -------------------------------------------
package services

import "github.com/onepiece-dz/drpc-server/services/api"

type Foo int

func (f Foo) Sum(args api.Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}
