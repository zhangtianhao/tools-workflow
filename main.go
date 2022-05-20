package main

import (
	"flag"
	"fmt"
	"math/big"
)

var convertErrItem = AlfredItem{Title: "转换错误", Subtitle: "检查要转换的数"}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
func has0bPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'b' || input[1] == 'B')
}
func main() {
	// 要转换的进制，可能值有：b（二进制）、o（八进制）、d（十进制）、x（十六进制）、
	var srcAry string
	// 要转换的数，这个数有可能会带上进制标识，比如：0x45ac
	var srcNum string

	flag.StringVar(&srcAry, "s", "", "要转换的数的进制")
	flag.StringVar(&srcNum, "n", "", "要转换的数")

	flag.Parse()
	var itemSlice = make([]AlfredItem, 0, 3)
	var num *big.Int
	switch srcAry {
	case "b":
		if has0bPrefix(srcNum) {
			srcNum = srcNum[2:]
		}
		num, _ = new(big.Int).SetString(srcNum, 2)
	case "o":
		num, _ = new(big.Int).SetString(srcNum, 8)
	case "d":
		num, _ = new(big.Int).SetString(srcNum, 10)
	case "x":
		if has0xPrefix(srcNum) {
			srcNum = srcNum[2:]
		}
		num, _ = new(big.Int).SetString(srcNum, 16)
	default:
		itemSlice = append(itemSlice, convertErrItem)
	}
	// 转十进制
	numStr := num.Text(10)
	itemSlice = append(itemSlice, AlfredItem{Title: numStr, Subtitle: "十进制", Arg: numStr})
	// 转十六进制
	numStr = num.Text(16)
	itemSlice = append(itemSlice, AlfredItem{Title: numStr, Subtitle: "十六进制", Arg: numStr})
	// 转二进制
	numStr = num.Text(2)
	itemSlice = append(itemSlice, AlfredItem{Title: numStr, Subtitle: "二进制", Arg: numStr})
	// 转八进制
	numStr = num.Text(8)
	itemSlice = append(itemSlice, AlfredItem{Title: numStr, Subtitle: "八进制", Arg: numStr})

	alfredResult := AlfredList{itemSlice}
	result := alfredResult.ToJson()
	fmt.Print(result)
}
