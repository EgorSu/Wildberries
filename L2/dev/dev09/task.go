package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

const (
	var nameFile = "" 
	var urls = ""
	var contin = false
)
func main() {

	if urls == "" {
		err := Download(nameFile, os.Args[len(os.Args)-1], contin, 1023)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		file, err := os.Open(urls)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			URL := scanner.Text()
			err = Download(nameFile, URL, contin, 1023)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func Download(nameFile, URL string, contin bool, size int64) error {
	var format string
	var firstByte int64
	var file *os.File

	client := &http.Client{}
	resp, err := client.Head(URL) 
	if err != nil {
		return err
	}
	sizeData := resp.ContentLength
	sizeC := (sizeData / size) + 1

	if val, ok := resp.Header["Content-Type"]; ok {
		format = strings.Split(strings.TrimSuffix(strings.Fields(val[0])[0], ";"), "/")[1] 
	}

	if contin {
		file, err = os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}
		firstByte = fileInfo.Size()
		defer file.Close()
	}

	isRead := make(chan struct{})
	isWrite := make(chan struct{})
	chData := make(chan []byte, sizeC)
	chErr := make(chan error)

	go func() {

		for {

			Range := fmt.Sprintf("bytes=%d-%d", firstByte, firstByte+size)

			req, err := http.NewRequest("GET", URL, nil)
			if err != nil {
				chErr <- err
				close(chErr)
				isRead <- struct{}{}
				return
			}

			req.Header.Add("Range", Range)

			resp, err = client.Do(req)
			if err != nil {
				chErr <- err
				close(chErr)
				isRead <- struct{}{}
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				chErr <- err
				close(chErr)
				isRead <- struct{}{}
				return
			}

			chData <- body
			if resp.ContentLength < size || sizeData == 0 {
				break
			}
			firstByte += size + 1
		}

		close(chErr)
		isRead <- struct{}{}
	}()

	go func() {
		if !contin {
			file, err = os.Create(fmt.Sprintf("%s.%s", nameFile, format))
			if err != nil {
				chErr <- err
				close(chErr)
				close(isRead)
				return
			}
			defer file.Close()
		}

		data := make([]byte, 0, size)
		writer := bufio.NewWriter(file)
		for {
			select {
			case data = <-chData:
				_, err := writer.WriteString(string(data))
				if err != nil {
					log.Println(err)
				}
				writer.Flush()
			case <-isRead:
				isWrite <- struct{}{}
				close(isRead)
			}
		}

	}()

	select {
	case err = <-chErr:
		return err
	case <-isWrite:
		close(isWrite)
		return nil
	}
}
