package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var isSleep bool

func main() {
	n := flag.Uint64("n", 1, "collatz target num. infinite loop if it was the this flag is 1 and l flag is true.")
	l := flag.Bool("l", false, "loop from 1 to the target. infinite loop if it was the this flag is true and n flag is 1.")
	s := flag.Bool("s", false, "0.5 seconds sleep every calculation.")
	flag.Parse()

	num := *n
	isSleep = *s
	loop := *l

	if num < 1 {
		fmt.Println("invalid value:", num)
		os.Exit(1)
	}

	if loop {
		if num == 1 {
			for i := uint64(1); ; i++ {
				loopCollatz(i)
			}
		} else {
			for i := uint64(1); i <= num; i++ {
				loopCollatz(i)
			}
		}
	} else {
		ret := collatz(num)
		retCheck(ret)
	}
}

func loopCollatz(i uint64) {
	ret := collatz(i)
	fmt.Println("----------")
	retCheck(ret)
}

func collatz(num uint64) (ret uint64) {
	fmt.Println(num)

	if isSleep {
		time.Sleep(time.Second / 2)
	}

	if num > 1 {
		if num%2 == 1 {
			ret = collatz(3*num + 1)
		} else {
			ret = collatz(num / 2)
		}
	} else {
		ret = num
	}

	return
}

func retCheck(ret uint64) {
	if ret != 1 {
		fmt.Println("invalid collatz return:", ret)
		os.Exit(1)
	}
}
