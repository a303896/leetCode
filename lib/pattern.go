package lib

import (
	"fmt"
	"regexp"
	"strconv"
)

func PatternTest(str string, pattern string) {
	f := func(s string) string {
		v,_ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok,_ := regexp.Match(pattern, []byte(str)); ok {
		fmt.Println("Match found!!!!")
	}

	re,_ := regexp.Compile(pattern)
	str1 := re.ReplaceAllString(str, "##.#")
	fmt.Println("ReplaceAllString:" + str1)
	str2 := re.ReplaceAllStringFunc(str, f)
	fmt.Println("ReplaceAllStringFunc:" + str2)
}
