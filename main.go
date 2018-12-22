package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		nt := strings.Replace(t, "_", " ", -1)
		nt = strings.Title(nt)
		nt = strings.Replace(nt, " ", "", -1)
		fmt.Println(nt)
	}
}
