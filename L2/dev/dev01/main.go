package task1

import (
	"fmt"

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
