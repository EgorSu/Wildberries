package main

import (
	"fmt"
	"io"
	"os"

	"github.com/beevik/ntp"
)

func timeNow() error {
	resp, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	}
	fmt.Println(resp.Time.Local())
	now, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	}
	fmt.Println(now)
	return nil
}

func main() {
	err := timeNow()
	if err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}
