package verbose

import (
	"fmt"
	"os"
	"strconv"
	"surge/cmdutil/common"
	"surge/cmdutil/table"
)

var Verbose bool

// LogHttpRequest 打印 HTTP 请求调试信息
func LogHttpRequest(context common.Response) {
	if Verbose {
		tableCols := []string{"Item", "Detail"}
		rows := [][]string{
			//{"X-KEY", surgeApiConfig.XKey},
			{"REQUEST_URL", context.Request.Url},
			{"METHOD", context.Request.Method},
			{"PAYLOAD", context.Request.Payload},
			{"StatusCode", strconv.Itoa(context.StatusCode)},
			{"RESPONSE", context.Context},
		}
		table.Gen(tableCols, rows)
	}
}

// Println 替代 fmp.Println 输出调试信息
func Println(a ...interface{}) (n int, err error) {
	if Verbose {
		fmt.Fprintln(os.Stdout, a...)
	}
	return
}

func PrintArray(args []string) {
	fmt.Println("Args is")
	for _, value := range args {
		fmt.Println("\t", value)
	}
}
