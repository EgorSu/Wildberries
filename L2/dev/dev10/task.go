package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	timeOut := flag.Duration("timeout", time.Duration(1), "timeout")
	args := flag.Args()
	host := args[0]
	port := args[1]
	var err error

	client := &http.Client{Timeout: *timeOut}
	_, err = client.Head(host + ":" + port)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	std := make([]string, 0, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "push" {
			body := make([]byte, 0, 0)
			for _, v := range std {
				body = append(body, []byte(v)...)
			}
			req, err := http.NewRequest("POST", host+":"+port, bytes.NewBuffer(body))
			if err != nil {
				log.Fatalln(err)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(resp)

			std = make([]string, 0, 0)
			continue
		}
		std = append(std, text)
	}
}
