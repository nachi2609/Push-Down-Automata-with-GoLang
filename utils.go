package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	//"fmt"
)

var (
	// custom logger; global variable, initialize by calling initLogging() in main()
	CLog *log.Logger
)

func initLogging() {
	var err error
	file := os.Stderr
	// file, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	CLog = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	//custom logger without any timestamps
	CLog = log.New(file, "", 0)
}

func check(e error) {
	// handle an error
	if e != nil {
		//fmt.Println("error ", e.Error())
		CLog.Println("error ", e.Error())
		panic(e)
	}
}

func readAllStdin() string {
	// read all lines from stdin and return as a string
	scanner := bufio.NewScanner(os.Stdin)
	buf := []string{}
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}
	check(scanner.Err())
	return strings.Join(buf, "\n")
}

func getCliArgs() []string {
	// return arguments on the command line (without the (self) filename of the executable)
	buf := []string{}
	for _, v := range os.Args[1:] {
		if v != "--" {
			buf = append(buf, v)
		}
	}
	return buf
}
