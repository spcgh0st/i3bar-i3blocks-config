package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"

)

func main() {

	content, _ := ioutil.ReadFile("/proc/loadavg")

	splitted := strings.Split(string(content), " ")

	avg_1  := splitted[0]
	avg_5  := splitted[1]
	avg_15 := splitted[2]

	avg_1_number, _ := strconv.ParseFloat(avg_1, 4)

	fmt.Printf("%s, %s, %s\n", avg_1, avg_5, avg_15)
	fmt.Printf("%s\n", avg_1)
	
	if avg_1_number > 6 {
		fmt.Println("#FF0000")
	}

}