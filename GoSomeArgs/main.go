package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	greeting("My name is %s!\n出席番号は%dです", "眠すぎ太郎", 17)
}

func greeting(format string, args ...interface{}) {
	for _, value := range args {
		index := strings.Index(format, "%")
		if index == -1 {
			panic("yurusan!")
			break
		}
		cont := ""
		switch format[index+1] {
		case 's':
			cont = value.(string)
		case 'd':
			cont = strconv.Itoa(value.(int))
		case 'f':
			cont = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		}
		format = format[:index] + cont + format[index+2:]
	}
	fmt.Println(format)
}
