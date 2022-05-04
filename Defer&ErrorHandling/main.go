package main

import (
	"fmt"
	"net"
)

type error interface {
	Error() string
}

func main() {
	/*
		Asserting the underlying struct type & getting more
		information from the struct fields
	*/
	// f, err := os.Open("tes.txt")
	// if err != nil {
	// 	//if the concrete type of err is *os.PathError then
	// 	//pErr will have the underlying value of err
	// 	if pErr, ok := err.(*os.PathError); ok {
	// 		fmt.Println("Fail to open file at path", pErr.Path)
	// 		return
	// 	}
	// 	fmt.Println("Generic error", err)
	// 	return
	// }
	// fmt.Println(f.Name(), "Open successfully")

	/*
		Asserting the underlying struct type & getting more
		information using method
	*/
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}
