package events

import (
	"time"

	"github.com/yuttasakcom/go-apis/middleware"
)

// ClearLastRequestsIPs event
func ClearLastRequestsIPs() {
	for {
		middleware.LastRequestsIPs = []string{}
		time.Sleep(time.Minute * 1)
	}
}

// ClearBlockedIPs event
func ClearBlockedIPs() {
	for {
		middleware.BlockedIPs = []string{}
		time.Sleep(time.Second * 5)
	}
}
