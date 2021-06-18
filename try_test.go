// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/6/18 - 5:37 下午 - UTC/GMT+08:00

// todo...

package try

import (
	"testing"
)

func TestDo(t *testing.T) {
	ex := new(DefaultException)
	Do(func() error {
		i := 0 / 1
		t.Log(i)
		return nil
	}, ex).Catch()
}
