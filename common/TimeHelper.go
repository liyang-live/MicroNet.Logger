package common

import "fmt"

/*
	将输入的date按照指定的partten格式化
	@version 1.0 目前只支持正常的格式，其余变态格式暂不支持
		如：yyyy,yyyy-MM,yyyy-MM-dd,yyyy-MM-dd HH,yyyy-MM-dd HH:mm,yyyy-MM-dd HH:mm:ss
	@date string 要被格式化的日期
	@partten stirng 指定格式
	@author wangdy
	return 返回格式化后的日期
*/
func TimFormat(date, partten string) string {
	var ptnLength = len(partten)
	var dateLength = len(date)
	returnStr := date
	switch partten {
	case partten:
		if ptnLength <= dateLength {
			return Substr(date, 0, ptnLength)
		} else {
			var chaLen = ptnLength - dateLength
			if dateLength == 4 {
				if chaLen == 3 {
					returnStr += "-01"
					return returnStr
				} else if chaLen == 6 {
					returnStr += "-01-01"
					return returnStr
				} else if chaLen == 9 {
					returnStr += "-01-01 00"
					return returnStr
				} else if chaLen == 12 {
					returnStr += "-01-01 00:00"
					return returnStr
				} else if chaLen == 15 {
					returnStr += "-01-01 00:00:00"
					return returnStr
				}
			} else if dateLength == 7 {
				if chaLen == 3 {
					returnStr += "-01"
					return returnStr
				} else if chaLen == 6 {
					returnStr += "-01 00"
					return returnStr
				} else if chaLen == 9 {
					returnStr += "-01 00:00"
					return returnStr
				} else if chaLen == 12 {
					returnStr += "-01 00:00:00"
					return returnStr
				}
			} else if dateLength == 10 {
				if chaLen == 3 {
					returnStr += " 00"
					return returnStr
				} else if chaLen == 6 {
					returnStr += " 00:00"
					return returnStr
				} else if chaLen == 9 {
					returnStr += " 00:00:00"
					return returnStr
				}
			} else if dateLength == 13 {
				if chaLen == 3 {
					returnStr += ":00"
					return returnStr
				} else if chaLen == 6 {
					returnStr += ":00:00"
					return returnStr
				}
			} else if dateLength == 16 && chaLen == 3 {
				returnStr += ":00"
				return returnStr
			}
			return "xxxxxxxxxx" + date
		}
	default:
		fmt.Println("it's default case!")
		return date
	}
}

/*
	按照格式截取字符串
	@str string 要被截取的源字符串
	@start int 开始索引，从0开始
	@length int 要截取的长度
	@author wangdy
	return 返回截取后字符串
*/
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
