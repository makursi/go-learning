//返回 error：写 divide(a, b int) (int, error)，b == 0 时返回错误，否则返回商。在 main 里分别用 (10, 2) 和 (10, 0) 调用并处理。

package main

import "errors"

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("除数不能为零")
	}

	return a / b, nil
}
