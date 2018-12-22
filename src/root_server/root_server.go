package root_server

import (
	"fmt"
	"strings"
)

func Root_server() {
	var ip string
	fmt.Scan(&ip)
	b := strings.Split(ip, ".")
	fmt.Printf("Length of split: %d\n", len(b))
}
