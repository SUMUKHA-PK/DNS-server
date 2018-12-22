package src

import (
	"fmt"
	"strings"
)

func src() {
	var ip string
	fmt.Scan(&ip)
	b := strings.Split(ip, ".")
	fmt.Printf("Length of split: %d", len(b))
	fmt.Printf("YO")
}

func main() {
	src()
}
