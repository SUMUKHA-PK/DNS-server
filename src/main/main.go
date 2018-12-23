package main

import (
	"fmt"
	"strings"

	"../dotcom/amazon"
	"../dotcom/dell"
	"../dotedu"
	"../dotedu/mit"
	"../dotedu/nitk"
	"../dotin"
	"../dotin/dotac"
	"../dotin/dotac/nitk/cse"
	"../dotin/dotac/nitk/ece"
	"../dotin/dotac/nitk/eee"
	"../root_server"
)

//                         STRUCTURE OF THE DNS SERVERS

// 						      root
// 							|
// 				---------------------------------------------------
// 				|                          |                      |
// 		             ------                     ------                 ------
// 			     |.in |                     |.com|                 |.edu|
// 			     ------                     ------                 ------
// 				|                          |                      |
// 			   ------------               -----------            -----------
// 			   |          |               |         |            |          |
// 		         -----      --------      --------    ------       ------      -----
// 		         |.ac|      |google|      |amazon|	 |dell|       |nitk|      |mit|
// 		         -----      --------      --------    ------       ------      -----
// 		          |            |            |           |            |           |
// 	                ------         -----          -----      -----        -----       -----
// 	               |nitk|  	      |www|          |www|      |www|        |www|       |web|
// 	               ------
// 	   	         |
//            ----------------------
// 	      |          |         |
//          ----        ----      ----
//         |cse|       |ece|     |eee|
//         -----       -----     -----

func main() {

	var ip string
	fmt.Scan(&ip)
	b := strings.Split(ip, ".")
	fmt.Printf("Length of split: %d\n", len(b))

	root_server.Root_server(b)
	fmt.Printf("|----")
	dotedu.DotEdu()
	fmt.Printf("     |----")
	nitk.Nitk()
	fmt.Printf("     |----")
	mit.Mit()
	dotin.DotIn()
	fmt.Printf("     |----")
	amazon.Amazon()
	fmt.Printf("     |----")
	dell.Dell()
	dotin.DotIn()
	fmt.Printf("     |----")
	dotac.DotAc()
	fmt.Printf("          |----")
	cse.Cse()
	fmt.Printf("          |----")
	ece.Ece()
	fmt.Printf("          |----")
	eee.Eee()

}
