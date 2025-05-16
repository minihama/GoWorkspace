package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindString("paranormal"))
	fmt.Println(re.FindString("paranormal"))
	fmt.Println(re.FindString("graal"))
	fmt.Println(re.FindString("none"))
	fmt.Println(re.FindStringIndex("paranormal"))
	fmt.Println(re.FindStringIndex("paranormal"))
	fmt.Println(re.FindStringIndex("graal"))
	fmt.Println(re.FindStringIndex("none"))
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 4))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", 4))
	fmt.Println(re.FindAllStringIndex("graal", -1))
	fmt.Println(re.FindAllStringIndex("none", -1))
}
