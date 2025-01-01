package ping

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Handler struct {
	Hosts          []string
	Sleeptime      int
	LastStatus     bool
	LastPingTime   time.Time
	LastPingResult []string
	LastProtocol   string
}

func NewHandler(hosts []string, sleep int) *Handler {
	return &Handler{Hosts: hosts, Sleeptime: sleep, LastStatus: true, LastProtocol: "Never"}
}

func (h *Handler) StartPinging() {
	result := make([]string, len(h.Hosts))
	for {
		isDown := true
		for index, host := range h.Hosts {
			cmd := exec.Command("ping", host, "-c", "1")
			_, err := cmd.Output()
			if err == nil {
				isDown = false
				result[index] = "UP"
			} else {
				result[index] = "Down"
			}
		}
		h.LastPingResult = result
		h.LastPingTime = time.Now()

		if isDown {
			if h.LastStatus {
				h.LastStatus = false
				log.Println("All ips are down. Following shutdown protocol")
				h.followProtocol(false)
			}
		} else {
			if !h.LastStatus {
				h.LastStatus = true
				log.Println("Ips are back up. Following normal protocol")
				h.followProtocol(true)
			}
		}
		time.Sleep(time.Duration(h.Sleeptime) * time.Second)
	}
}

func (h *Handler) followProtocol(status bool) {
	path := ""
	if status {
		path = os.Getenv("SCRIPT_UP")
	} else {
		path = os.Getenv("SCRIPT_DOWN")
	}
	cmd := exec.Command("bash", path)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("**\nScript error: %s\n%s\n***\n", path, err)
		return
	}
	log.Printf("**\nProtocol output: %s\n%s\n***\n", path, out)
	h.LastProtocol = fmt.Sprintf("%t: %s", status, time.Now().Format("15:04:05 02.01.2006"))
}
