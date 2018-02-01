package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// LIMIT request
const LIMIT = 2

// LastRequestsIPs []string
var LastRequestsIPs []string

// BlockedIPs []string
var BlockedIPs []string

func existsLastRequest(ipAddr string) bool {
	for _, ip := range LastRequestsIPs {
		if ip == ipAddr {
			return true
		}
	}
	return false
}

func existsBlockedIP(ipAddr string) bool {
	for _, ip := range BlockedIPs {
		if ip == ipAddr {
			return true
		}
	}
	return false
}

// RateLimit http.Handler
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ipAddr := strings.Split(r.RemoteAddr, ":")[0]

		if existsBlockedIP(ipAddr) {
			fmt.Printf("BlockedIPs: %v\n", BlockedIPs)
			http.Error(w, "block ip", http.StatusTooManyRequests)
			return
		}

		// how many requests the current IP made in last 5 mins
		requestCounter := 0

		for _, ip := range LastRequestsIPs {
			if ip == ipAddr {
				requestCounter++
			}
		}

		if requestCounter >= LIMIT {
			BlockedIPs = append(BlockedIPs, ipAddr)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		LastRequestsIPs = append(LastRequestsIPs, ipAddr)

		// Don't cut the chain of middlewares
		if next == nil {
			http.DefaultServeMux.ServeHTTP(w, r)
			return
		}

		fmt.Printf("LastRequestsIPs: %v\n", LastRequestsIPs)

		next.ServeHTTP(w, r)
	})
}
