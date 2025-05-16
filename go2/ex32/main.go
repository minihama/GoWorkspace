package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	matches := re.FindStringSubmatch("Alan Turing")
	fmt.Printf("%q\n", re.SubexpNames())
	lastIndex := re.SubexpIndex("last")
	fmt.Printf("last => %d\n", lastIndex)
	fmt.Println(matches[lastIndex])
}
