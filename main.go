package main

import (
	"fmt"
	"net"
	"sort"
	"github.com/schollz/progressbar"
	lol "github.com/kris-nova/lolgopher"
	"os"
)

var target string
var p string

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf(target + ":%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {

	titleArt := `
.▄▄ · ▪  • ▌ ▄ ·.  ▄▄▄·▄▄▌  ▄▄▄ .▄▄▄▄▄ ▄ .▄▄▄▄  ▄▄▄ . ▄▄▄· ·▄▄▄▄  ▄▄▄ .·▄▄▄▄  
▐█ ▀. ██ ·██ ▐███▪▐█ ▄███•  ▀▄.▀·•██  ██▪▐█▀▄ █·▀▄.▀·▐█ ▀█ ██▪ ██ ▀▄.▀·██▪ ██ 
▄▀▀▀█▄▐█·▐█ ▌▐▌▐█· ██▀·██▪  ▐▀▀▪▄ ▐█.▪██▀▐█▐▀▀▄ ▐▀▀▪▄▄█▀▀█ ▐█· ▐█▌▐▀▀▪▄▐█· ▐█▌
▐█▄▪▐█▐█▌██ ██▌▐█▌▐█▪·•▐█▌▐▌▐█▄▄▌ ▐█▌·██▌▐▀▐█•█▌▐█▄▄▌▐█ ▪▐▌██. ██ ▐█▄▄▌██. ██ 
 ▀▀▀▀ ▀▀▀▀▀  █▪▀▀▀.▀   .▀▀▀  ▀▀▀  ▀▀▀ ▀▀▀ ·.▀  ▀ ▀▀▀  ▀  ▀ ▀▀▀▀▀•  ▀▀▀ ▀▀▀▀▀• 
 ▄▄▄·      ▄▄▄  ▄▄▄▄▄    .▄▄ ·  ▄▄·  ▄▄▄·  ▐ ▄  ▐ ▄ ▄▄▄ .▄▄▄                  
▐█ ▄█▪     ▀▄ █·•██      ▐█ ▀. ▐█ ▌▪▐█ ▀█ •█▌▐█•█▌▐█▀▄.▀·▀▄ █·                
 ██▀· ▄█▀▄ ▐▀▀▄  ▐█.▪    ▄▀▀▀█▄██ ▄▄▄█▀▀█ ▐█▐▐▌▐█▐▐▌▐▀▀▪▄▐▀▀▄                 
▐█▪·•▐█▌.▐▌▐█•█▌ ▐█▌·    ▐█▄▪▐█▐███▌▐█ ▪▐▌██▐█▌██▐█▌▐█▄▄▌▐█•█▌                
.▀    ▀█▄▀▪.▀  ▀ ▀▀▀      ▀▀▀▀ ·▀▀▀  ▀  ▀ ▀▀ █▪▀▀ █▪ ▀▀▀ .▀  ▀                
 `

lol.Println(titleArt)
	fmt.Println("")
	lol.Println(" ==============================================")
	lol.Println(" |  Simple MultiThread Port Scanner - Golang  |")
	lol.Println(" |             by Insolent-M1nx               |")
	lol.Println(" ==============================================")
	fmt.Println("")

	//user
	fmt.Println("")
	lol.Printf(" Please Enter Target To Port Scan : ")
	_, err := fmt.Scanln(&target)
	if err != nil {
		fmt.Print(os.Stderr, err)
	}
	
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	fmt.Println("")
	lol.Println(" Scanning target for open ports...")
	lol.Println("")
	bar := progressbar.Default(1024)
	for i := 0; i < 1024; i++ {
		bar.Add(1)
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
		fmt.Println("")
		lol.Println(" Port   Status")
		lol.Println(" -------------")
	for _, port := range openports {

		lol.Printf(" %d      OPEN \n", port)
	}
}