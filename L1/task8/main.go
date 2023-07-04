package main

import (
	"flag"
	"fmt"
)

func main() {
	//параметры получаем через флаги (как в 4 задании)
	val := flag.Int64("value", 1, "value for change")
	bitNum := flag.Int("number", 0, "number of bit for change")
	bitVal := flag.Int("bit", 0, "new value of a bit: true == 1, false == 0")
	flag.Parse()

	if *bitNum > 64 {
		fmt.Println("The bit number can not be greater than 64")
		return
	}

	fmt.Printf("value before: \t%-08b\n", *val)
	//для изменения заначения бита создадим число с одним битом == 1(тот бит который будем менять)
	//для выставления зачения == 1 используем поразрядную дизъюнкцию (всегда вернёт 1 т.к. наш бит == 1)
	//для выставления значения == 0 используем сброс бита
	switch *bitVal {
	case 1:
		*val = *val | 1<<*bitNum
	case 0:
		*val = *val &^ (1 << *bitNum)
	default:
		fmt.Println("The bit value can be only one or zero")
	}

	fmt.Printf("value after: \t%08b\n", *val)
}
