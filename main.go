package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ozcanarican/network-watcher/internal/ping"
)

func main() {
	godotenv.Load()
	hText := os.Getenv("HOSTS")
	hText = strings.ReplaceAll(hText, " ", "")
	hosts := strings.Split(hText, ",")
	fmt.Println("---------------------")
	print(fmt.Sprintf("%d host has loaded: %s", len(hosts), hosts))
	print("Pinger has started to checking ip status")
	print("Available commands: status, exit, clear")
	reader := bufio.NewReader(os.Stdin)
	pinger := ping.NewHandler(hosts, 10)
	go pinger.StartPinging()
	for {
		text, _ := reader.ReadString('\n')
		cmd := strings.ReplaceAll(text, "\n", "")
		fmt.Println("---------------------")
		if cmd == "exit" {
			break
		} else if cmd == "status" {
			list := ""
			for index, st := range pinger.LastPingResult {
				list += fmt.Sprintf("%s: %s\n", pinger.Hosts[index], st)
			}
			print(fmt.Sprintf("IP health status is: %t\nLast check time was on %s\n--\n%s--\nProtocol script runned on: %s", pinger.LastStatus, pinger.LastPingTime.Format("02.01.2006 15:04:05"), list, pinger.LastProtocol))
		} else if cmd == "clear" {
			fmt.Print("\033[H\033[2J")
		} else {
			print(fmt.Sprintf("Unknown command: %s", cmd))
		}
	}
}

func print(s string) {
	fmt.Println(s)
	fmt.Println("---------------------")
}
