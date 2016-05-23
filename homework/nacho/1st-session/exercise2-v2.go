package main

import (
	"fmt"
	"strings"
)

var full_router_name string = "dr02.atl21.net.google.com"
var router_name string = full_router_name[0:10]
var role string = full_router_name[0:2]
var metro string = strings.ToUpper(full_router_name[5:8])
var pop string = full_router_name[4:10]

func pair_router(hostname string){
	dr := string(hostname[3])
	var dr1,dr2,dr3,dr4 string = "1","2","3","4" // declaring the DR router numbers to give flexibility if they change.
	if dr == dr1 {
		fmt.Println(role + "02" + pop)
	} else if dr == dr3 {
		fmt.Println(role + "04" + pop)
	} else if dr == dr2 {
		fmt.Println(role + "01" + pop)
	} else if dr == dr4 {
		fmt.Println(role + "03" + pop)
	}
}

func first_router(pop string){
	fmt.Println(role + "01." + pop)
}

func main() {
	fmt.Println(full_router_name == "dr04.atl21.net.google.com") // print true if router name is “dr04.atl21.net.google.com”, otherwise print false.
	fmt.Println(role == "dr") // print true if the router is a DR, otherwise print false.
	fmt.Println(strings.ToLower(router_name)) // print dr02.atl21
	fmt.Println(strings.ToUpper(router_name)) // print DR02.ATL21
	fmt.Println(metro == "ATL") // print true if the device is in ATL metro, otherwise print false.
	pair_router(full_router_name) // Calling the func to print the pair router in the metro. If dr01 return dr02 and opposite, same if dr03 <-> dr04
	first_router("atl11") // Calling the func to print first router in pop atl11 as argument.
}
