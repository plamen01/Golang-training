package main

import (
	"fmt"
	"strings"
)

var router_name string = "dr02.atl21.net.google.com"
var short_router_name_end int = 10
var short_router_name string = router_name[0:short_router_name_end]
var role string = router_name[0:2]
var metro string = strings.ToUpper(router_name[5:8])
var pop string = router_name[4:short_router_name_end]

func pair_router(hostname string){
	number := string(hostname[3])
	var first,second,third,fourth string = "1","2","3","4"
	if number == first {
		fmt.Println(role + "02" + pop)
	} else if number == third {
		fmt.Println(role + "04" + pop)
	} else if number == second {
		fmt.Println(role + "01" + pop)
	} else if number == fourth {
		fmt.Println(role + "03" + pop)
	}
}

func first_router(pop string){
	fmt.Println(role + "01." + pop)
}

func main() {
	fmt.Println(router_name == "dr04.atl21.net.google.com") // print true if router name is “dr04.atl21.net.google.com”, otherwise print false.
	fmt.Println(role == "dr") // print true if the router is a DR, otherwise print false.
	fmt.Println(strings.ToLower(short_router_name)) // print dr02.atl21
	fmt.Println(strings.ToUpper(short_router_name)) // print DR02.ATL21
	fmt.Println(metro == "ATL") // print true if the device is in ATL metro, otherwise print false.
	pair_router(router_name) // Calling the func to print the pair router in the metro. If dr01 return dr02 and opposite, same if dr03 <-> dr04
	first_router("atl11") // Calling the func to print first router in pop atl11 as argument.
}
