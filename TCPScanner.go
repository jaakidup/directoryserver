package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// TCPScanner ...
// Ports to scan can be from 0 to 65535
func TCPScanner(host string, startPort int, endPort int, timeout time.Duration) []int {
	ports := []int{}
	waitGroup := &sync.WaitGroup{}
	mutex := sync.Mutex{}
	for port := startPort; port <= endPort; port++ {
		waitGroup.Add(1)

		go func(host string, port int, timeout time.Duration) {
			if OpenPort(host, port, timeout) {
				mutex.Lock()
				ports = append(ports, port)
				mutex.Unlock()
			}
			waitGroup.Done()
		}(host, port, timeout)

	}
	waitGroup.Wait()
	return ports
}

// OpenPort ...
func OpenPort(host string, port int, timeout time.Duration) bool {
	connString := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", connString, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
