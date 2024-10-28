package main

import (
	"StringUnpack/unpack"
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	fmt.Println(unpack.Pack("\n\n\n\n")) // \n4
	input := flag.String("input", "", "input for Unpack/Pack")
	inputRaw := flag.String("raw", "", "input for Unpack Raw")
	isDaemon := flag.Bool("daemon", false, "run in daemon mode")
	pack := flag.Bool("pack", false, "run in pack mode")
	flag.Parse()
	if *isDaemon {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go readInput(*pack)
		for sig := range sigs {
			fmt.Printf("\nReceived signal: %s. Exiting...\n", sig)
			return
		}
	}
	if *input != "" {
		if *pack {
			res, err := unpack.Pack(*input)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(res)
			}
			return
		}
		//?d1\n4
		fmt.Println(*input)
		//?d\n\n\n\n
		fmt.Println(unpack.Unpack("d1\n4", false))
		//?returns like "dnnnnn"
		res, err := unpack.Unpack(*input, false)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
	if *inputRaw != "" {
		res, err := unpack.Unpack(*inputRaw, true)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}

func readInput(pack bool) {
	if pack {
		processPack()
	} else {
		processUnpack()
	}
}

func processPack() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You entered pack mode. Enter [STRING] (type ctrl+c to quit):")
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if len(input) != 1 {
			fmt.Println("Wrong input. Enter [STRING] (type ctrl+c to quit):")
			continue
		}
		// Process the input
		res, err := unpack.Pack(input[0])
		fmt.Println("result:")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

func processUnpack() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You enetered unpack mode. Enter [STRING ISRAW] (type ctrl+c to quit):")
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if len(input) != 2 {
			fmt.Println("Wrong input. Enter [STRING ISRAW] (type ctrl+c to quit):")
			continue
		}
		isRaw, _ := strconv.ParseBool(input[1])
		// Process the input
		res, err := unpack.Unpack(input[0], isRaw)
		fmt.Println("result:")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}
