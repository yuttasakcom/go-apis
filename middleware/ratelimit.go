package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// LIMIT request
const LIMIT = 60

var lastRequestsIPs []string

var blockedIPs []string

func existsBlockedIP(ipAddr string) bool {
	for _, ip := range blockedIPs {
		if ip == ipAddr {
			return true
		}
	}
	return false
}

func existsLastRequest(ipAddr string) bool {
	for _, ip := range lastRequestsIPs {
		if ip == ipAddr {
			return true
		}
	}
	return false
}

// Clears lastRequestsIPs array every 5 mins
func clearLastRequestsIPs() {
	for {
		lastRequestsIPs = []string{}
		time.Sleep(time.Minute * 5)
	}
}

// Clears blockedIPs array every 6 hours
func clearBlockedIPs() {
	for {
		blockedIPs = []string{}
		time.Sleep(time.Hour * 6)
	}
}

// RateLimit http.Handler
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, "RateLimit")

		ipAddr := strings.Split(r.RemoteAddr, ":")[0]

		if existsBlockedIP(ipAddr) {
			http.Error(w, "block ip", http.StatusTooManyRequests)
			return
		}

		// how many requests the current IP made in last 5 mins
		requestCounter := 0

		for _, ip := range lastRequestsIPs {
			if ip == ipAddr {
				requestCounter++
			}
		}

		if requestCounter >= LIMIT {
			blockedIPs = append(blockedIPs, ipAddr)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		lastRequestsIPs = append(lastRequestsIPs, ipAddr)

		// Don't cut the chain of middlewares
		if next == nil {
			http.DefaultServeMux.ServeHTTP(w, r)
			return
		}

		fmt.Printf("lastRequestsIPs: %v, blockedIPs: %v", lastRequestsIPs, blockedIPs)

		next.ServeHTTP(w, r)
	})
}
