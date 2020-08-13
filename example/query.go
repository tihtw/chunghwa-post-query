package main

// 手機條碼驗證

import (
	"fmt"
	"github.com/tihtw/chunghwa-post-query"
	"os"
)

func main() {
	fmt.Println("start")

	host := os.Getenv("INVOICE_CONNECTION_HOST")

	c := cpquery.NewClient()
	// c.ApiKey = apiKey

	if host != "" {
		c.ConnectionHost = host
	}
	banUnitTpStatus, err := c.QueryMailDetail("44261122507318")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("status:", banUnitTpStatus)
}
