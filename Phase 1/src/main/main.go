package main

import (
	"fmt"

	"../dotcom"
	"../dotcom/amazon"
	"../dotcom/dell"
	"../dotedu"
	"../dotedu/mit"
	"../dotedu/nitk"
	"../dotin"
	"../dotin/dotac"
	"../dotin/dotac/nitkac"
	"../dotin/google"
	"../root_server"
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

	var IP_List_Name_root = []string{"dotin", "dotcom", "dotedu"}
	var IP_List_Addr_root = []string{"127.0.1.0", "127.0.1.1", "127.0.1.2"}
	var IP_List_Name_com = []string{"amazon", "dell"}
	var IP_List_Addr_com = []string{"127.0.2.2", "127.0.2.3"}
	var IP_List_Name_in = []string{"dotac", "google"}
	var IP_List_Addr_in = []string{"127.0.2.0", "127.0.2.1"}
	var IP_List_Name_ac = []string{"nitkac"}
	var IP_List_Addr_ac = []string{"127.0.3.0"}
	var IP_List_Name_edu = []string{"nitk", "mit"}
	var IP_List_Addr_edu = []string{"127.0.2.4", "127.0.2.5"}

	go func() {
		root_server.Root_server(IP_List_Name_root, IP_List_Addr_root)
	}()
	go func() {
		dotin.DotIn(IP_List_Name_in, IP_List_Addr_in)
	}()
	go func() {
		dotcom.DotCom(IP_List_Name_com, IP_List_Addr_com)
	}()
	go func() {
		dotedu.DotEdu(IP_List_Name_edu, IP_List_Addr_edu)
	}()
	go func() {
		dotac.DotAc(IP_List_Name_ac, IP_List_Addr_ac)
	}()
	go func() {
		amazon.Amazon(IP_List_Name_com, IP_List_Addr_com)
	}()
	go func() {
		dell.Dell(IP_List_Name_com, IP_List_Name_com)
	}()
	go func() {
		nitk.Nitk(IP_List_Name_edu, IP_List_Addr_edu)
	}()
	go func() {
		mit.Mit(IP_List_Name_edu, IP_List_Addr_edu)
	}()
	go func() {
		google.Google(IP_List_Name_in, IP_List_Addr_in)
	}()
	go func() {
		nitkac.Nitkac(IP_List_Name_in, IP_List_Addr_in)
	}()

	fmt.Scanln()
	fmt.Println("done")

}
