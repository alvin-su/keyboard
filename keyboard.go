//Package 从键盘获取用户的输入.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)
// GetFloat 将键盘中用户的输入去掉空格并转为float返回.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
