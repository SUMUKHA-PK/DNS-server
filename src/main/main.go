package main

import (
	"dotcom/amazon"
	"dotcom/dell"
	"dotedu"
	"dotedu/mit"
	"dotedu/nitk"
	"dotin"
	"dotin/dotac"
	"dotin/dotac/nitk/cse"
	"dotin/dotac/nitk/ece"
	"dotin/dotac/nitk/eee"
	"fmt"
	"root_server"
)

//                         STRUCTURE OF THE DNS SERVERS

// 						      root (127.0.0.1)
// 							|
// 				---------------------------------------------------
// 				|                          |                      |
// 		             ------                     ------                 ------
// 			     |.in | (127.0.1.0)         |.com| (127.0.1.1)     |.edu| (127.0.1.2)
// 			     ------                     ------                 ------
// 				|                          |                      |
// 			   ------------               -----------            -----------
// 			   |          |               |         |            |          |
// 		         -----      --------      --------    ------       ------      -----
// 		         |.ac| (2.0)|google|(2.1) |amazon|(2.2)|dell|(2.3)|nitk|(2.4) |mit|(2.5)
// 		         -----      --------      --------    ------       ------      -----
// 		          |            |            |           |            |           |
// 	                ------         -----          -----      -----        -----       -----
// 	               |nitk|(3.0)    |www|          |www|      |www|        |www|       |web|
// 	               ------
// 	   	         |
//            ----------------------
// 	      |          |         |
//          ----        ----      ----
//         |cse|       |ece|     |eee|
//         -----       -----     -----

func main() {

	fmt.Printf("I am the Root server!\n")
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
	root_server.Root_server()
}
