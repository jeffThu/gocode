package main

import "fmt"

func main() {
	fmt.Println(101, numberToChinese(101))
	fmt.Println(1012, numberToChinese(1012))
	fmt.Println(100, numberToChinese(100))

	fmt.Println(1000000, numberToChinese(1000000))
	fmt.Println(10010, numberToChinese(10010))

	fmt.Println(210100, numberToChinese(210100))
	fmt.Println(11, numberToChinese(11)) //x
}

// 阿拉伯数字转汉字
func numberToChinese(n int) (chinese string) {
	base := 10000
	items := []string{
		"", "万", "亿",
	}
	for i := 0; n > 0; i++ {
		numb4 := n % base
		n /= base
		chinese = numb4ToChinese(numb4) + items[i] + chinese
	}
	runes := []rune(chinese)
	if runes[0] == '零' {
		runes = runes[1:]
	}
	chinese = string(runes)
	return
}

func numb4ToChinese(n int) (chinese string) {
	if n >= 10000 {
		return
	}
	var s stack
	base := 10
	for n > 0 {
		v := n % base
		n /= base
		s = append(s, v)
	}
	for len(s) < 4 { //不够四位，补前缀零
		s = append(s, 0)
	}
	values := []string{
		"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十",
	}
	items := []string{
		"十", "百", "千",
	}
	for len(s) > 0 {
		l := len(s)
		value := s.pop()
		//如果是最后一个零了,并且下一个还有非零数，则加“零”，否则不加任何
		if value == 0 {
			ok, top := s.top()
			if ok == true && top != 0 {
				chinese += values[value]
			}
			continue
		}
		var item string
		if l-2 >= 0 {
			item = items[l-2]
		}
		chinese = chinese + values[value] + item
	}

	return
}

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}
func (s *stack) pop() (v int) {
	if len(*s) <= 0 {
		return
	}
	v = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}
func (s stack) top() (ok bool, v int) {
	if len(s) == 0 {
		return
	}
	ok, v = true, s[len(s)-1]
	return
}
