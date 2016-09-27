package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getTransmissionData(outPut bytes.Buffer) {
	var ll []string
	nl := 0
	for {
		nl++
		line, err := outPut.ReadString('\n')
		if err != nil {
			break
		}
		ll = strings.Fields(line)
		if nl == 5 {
			fmt.Printf("%v,%v,%v\n", time.Now().Format("2006-01-02 15:04:05"), ll[1], ll[2])
		}
	}
}

func getPingData(outPut bytes.Buffer) {
	var ll []string
	nl := 0
	for {
		nl++
		line, err := outPut.ReadString('\n')
		if err != nil {
			break
		}
		ll = strings.Fields(line)
		if nl == 3 {
			bytes := strings.Split(ll[3:][0], "=")[1]
			timeInMS := strings.Split(ll[3:][1], "=")[1]
			responseTime := strings.Split(timeInMS, "ms")[0]
			ttl := strings.Split(ll[3:][2], "=")[1]
			fmt.Printf("%v,%v,%v,%v\n", time.Now().Format("2006-01-02 15:04:05"), bytes, responseTime, ttl)
		}
	}
}

func main() {
	transmission := flag.Bool("transmission", false, "Every second printing local network transmission, format: timestamp,received,sent")
	ping := flag.String("ping", "", "Every second printing the latency from remote address, format: timestamp,bytes,time,ttl")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *transmission {
		for {
			cmd := exec.Command("netstat", "-e")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			getTransmissionData(out)
			time.Sleep(time.Second * 1)
		}
	}

	if *ping != "" {
		for {
			cmd := exec.Command("ping", "-n", "1", *ping)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			getPingData(out)
			time.Sleep(time.Second * 1)
		}
	}
}
