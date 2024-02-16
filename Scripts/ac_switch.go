package main

import(
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func main() {

	r := regexp.MustCompile(`state:[ ]*(.*)`)

	out, err := exec.Command("/usr/bin/upower", "-i", "/org/freedesktop/UPower/devices/battery_BAT0").Output()

	if err != nil {
		log.Fatal(err)
	}

	state := r.FindStringSubmatch(string(out))[1]

	if state == "charging" {
		fmt.Println("")

 	}  else {
		fmt.Println("")
	}

}
