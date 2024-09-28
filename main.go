// main.go

package main

import (
	"flag"
	"fmt"
	"go-socks/socks5proxy"
	"log"
	"os"
)

func main() {
	var port uint
	flag.UintVar(&port, "port", 0, "Port to listen on")
	flag.Parse()

	if port == 0 || port > 65535 {
		fmt.Printf("Invalid port number: %d\nUsage: %s -port <port>\n", port, os.Args[0])
		os.Exit(1)
	}

	proxy := socks5proxy.NewSocks5Proxy()
	err := proxy.Open(uint16(port))
	if err != nil {
		log.Fatal("main: error opening proxy on port ", port, ": ", err.Error())
	}

	go func() {
		err = proxy.Serve()
		if err != nil {
			log.Fatal("main: error in proxy's serve: ", err.Error())
		}
	}()

	//ticker := time.NewTicker(time.Second)
	//mbDivider := float64(1024 * 1024)
	//for range ticker.C {
	//	result := proxy.Stats()
	//	var sumAverageReadSpeedMB float64
	//	var sumAverageWriteSpeedMB float64
	//	var sumMomentReadSpeedMB float64
	//	var sumMomentWriteSpeedMB float64
	//	for _, v := range result {
	//		now := time.Now()
	//		client := v.Client
	//		averageReadSpeedMB := float64(client.ReadBytes) / now.Sub(client.StartTime).Seconds() / mbDivider
	//		averageWriteSpeedMB := float64(client.WroteBytes) / now.Sub(client.StartTime).Seconds() / mbDivider
	//		momentReadSpeedMB := float64(client.ReadSinceLastStats) / now.Sub(client.LastStatsTime).Seconds() / mbDivider
	//		momentWriteSpeedMB := float64(client.WroteSinceLastStats) / now.Sub(client.LastStatsTime).Seconds() / mbDivider
	//		sumAverageReadSpeedMB += averageReadSpeedMB
	//		sumAverageWriteSpeedMB += averageWriteSpeedMB
	//		sumMomentReadSpeedMB += momentReadSpeedMB
	//		sumMomentWriteSpeedMB += momentWriteSpeedMB
	//	}
	//	fmt.Printf("average: read - %.2f MB/s, write - %.2f MB/s, moment: read - %.2f MB/s, write - %.2f MB/s\n",
	//		sumAverageReadSpeedMB, sumAverageWriteSpeedMB, sumMomentReadSpeedMB, sumMomentWriteSpeedMB)
	//}

	var str string
	_, _ = fmt.Scan(&str)
	err = proxy.Close()
	if err != nil {
		log.Fatal("main: error in proxy's close: ", err.Error())
	}
}
