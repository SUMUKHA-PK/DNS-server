package main

import (
	"dotcom"
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

	var IP_List_Name = []string{"dotin", "dotcom", "dotedu", "dotac", "google", "amazon", "dell", "nitk", "mit", "nitkac"}
	var IP_List_Addr = []string{"127.0.1.0", "127.0.1.1", "127.0.1.2", "127.0.2.0", "127.0.2.1", "127.0.2.2", "127.0.2.3", "127.0.2.4", "127.0.2.5", "127.0.3.0"}

	go func() {
		root_server.Root_server(IP_List_Name, IP_List_Addr)
	}()
	fmt.Printf("|----")
	go func() {
		dotedu.DotEdu()
	}()
	fmt.Printf("     |----")
	go func() {
		nitk.Nitk()
	}()
	fmt.Printf("     |----")
	go func() {
		mit.Mit()
	}()
	go func() {
		dotcom.DotCom(IP_List_Name, IP_List_Addr)
	}()
	fmt.Printf("     |----")
	go func() {
		amazon.Amazon(IP_List_Name, IP_List_Addr)
	}()
	fmt.Printf("     |----")
	go func() {
		dell.Dell()
	}()
	go func() {
		dotin.DotIn()
	}()
	fmt.Printf("     |----")
	go func() {
		dotac.DotAc()
	}()
	fmt.Printf("          |----")
	go func() {
		cse.Cse()
	}()
	fmt.Printf("          |----")
	go func() {
		ece.Ece()
	}()
	fmt.Printf("          |----")
	go func() {
		eee.Eee()
	}()

	fmt.Scanln()
	fmt.Println("done")

}
