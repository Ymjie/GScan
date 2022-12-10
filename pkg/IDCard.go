package pkg

import (
	"strconv"
	"strings"
	"time"
)

func IsIDCard(id string) bool {
	// 身份证位数不对
	if len(id) != 15 && len(id) != 18 {
		return false
	}

	// 转大写
	id = strings.ToUpper(id)

	if len(id) == 18 {
		// 验证算法
		if !checkValidNo18(id) {
			//fmt.Println(id, "身份证算法验证失败！")
			return false
		}

	} else {
		// 转18位
		id = idCard15To18(id)
	}

	// 生日验证
	if !checkBirthdayCode(id[6:14]) {
		//fmt.Println(id, "生日验证失败！")
		return false
	}

	// 验证地址
	//if !checkAddressCode(id[:6]) {
	//	fmt.Println(id, "地址验证失败！")
	//	return false
	//}
	//
	return true
}

//15位身份证转为18位
var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

// 15位转18位
func idCard15To18(id15 string) string {
	nLen := len(id15)
	if nLen != 15 {
		return "身份证不是15位！"
	}
	id18 := make([]byte, 0)
	id18 = append(id18, id15[:6]...)
	id18 = append(id18, '1', '9')
	id18 = append(id18, id15[6:]...)

	sum := 0
	for i, v := range id18 {
		n, _ := strconv.Atoi(string(v))
		sum += n * weight[i]
	}
	mod := sum % 11
	id18 = append(id18, validValue[mod])
	return string(id18)
}

//18位身份证校验码
func checkValidNo18(id string) bool {
	//string -> []byte
	id18 := []byte(id)
	nSum := 0
	for i := 0; i < len(id18)-1; i++ {
		n, _ := strconv.Atoi(string(id18[i]))
		nSum += n * weight[i]
	}
	//mod得出18位身份证校验码
	mod := nSum % 11
	if validValue[mod] == id18[17] {
		return true
	}

	return false
}

// 验证生日
func checkBirthdayCode(birthday string) bool {
	year, _ := strconv.Atoi(birthday[:4])
	month, _ := strconv.Atoi(birthday[4:6])
	day, _ := strconv.Atoi(birthday[6:])

	curYear, curMonth, curDay := time.Now().Date()
	//出生日期大于现在的日期
	if year < 1900 || year > curYear || month <= 0 || month > 12 || day <= 0 || day > 31 {
		return false
	}

	if year == curYear {
		if month > int(curMonth) {
			return false
		} else if month == int(curMonth) && day > curDay {
			return false
		}
	}

	//出生日期在2月份
	if 2 == month {
		//闰年2月只有29号
		if isLeapYear(year) && day > 29 {
			return false
		} else if day > 28 { //非闰年2月只有28号
			return false
		}
	} else if 4 == month || 6 == month || 9 == month || 11 == month { //小月只有30号
		if day > 30 {
			return false
		}
	}

	return true
}

// 判断是否为闰年
func isLeapYear(year int) bool {
	if year <= 0 {
		return false
	}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

// 验证地区
// strict: true-验证详细， false-验证省
//func checkAddressCode(address string) bool {
//	if _, ok := area[address]; ok {
//		return true
//	}
//
//	return false
//}
