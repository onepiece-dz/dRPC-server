// Package services -----------------------------
// @author   : deng zhi
// @time     : 2024/9/10 15:57
//
// -------------------------------------------
package services

import "time"

type Bar int

func (b Bar) Timeout(argv int, reply *int) error {
	time.Sleep(time.Second * 2)
	return nil
}
