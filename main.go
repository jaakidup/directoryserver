package main

import (
	"fmt"
	"net/http"
)

func init() {

}

// GetAdresses ...
func GetAdresses() []string {

	fmt.Println("Directory Server")
	ipAddresses, err := ExternalIPAdresses()
	if err != nil {
		panic(err)
	}

	return ipAddresses

	// fmt.Println("Discovered IPAdresses")
	// for _, ip := range ipAddresses {
	// 	fmt.Println(ip)
	// 	openPorts := TCPScanner(ip, 0, 65535, 1*time.Second)
	// 	fmt.Println("Ports: ", openPorts)
	// 	// hosts = append(hosts, )
	// }
}

// DoUserConfig ...
func DoUserConfig() {

	var host string
	var port string

	fmt.Print("Enter Host: ")
	fmt.Scanln(&host)
	if host == "" {
		host = "0.0.0.0"
	}

	fmt.Print("Entre port number: ")
	fmt.Scanln(&port)

	if port == "" {
		port = "8000"
	}

	adress := host + ":" + port

	fmt.Println(adress)

}

func main() {

	wait := make(chan bool)

	http.Handle("/", http.FileServer(http.Dir(".")))

	for _, adress := range GetAdresses() {
		fmt.Printf("Starting directory server on: %v\n", adress+":8000")
		go http.ListenAndServe(adress+":8000", nil)
	}
	fmt.Printf("Starting directory server on: %v\n", "127.0.0.1:8000")
	go http.ListenAndServe("127.0.0.1:8000", nil)

	<-wait

}
