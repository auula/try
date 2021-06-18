// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/6/18 - 5:37 下午 - UTC/GMT+08:00

// todo...

package try

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	Do(func() error {
		return errors.New("1111")
	}).Catch(func(exception DefaultException) {
		aa
	})
}
