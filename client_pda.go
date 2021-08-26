/*
# Execute as
#
$go run client_pda.go pda.go collection_functions.go utils.go -- data/pda1_spec.txt < data/pda1_inp01.txt
#
# or as
#
$go build client_pda.go pda.go collection_functions.go utils.go
$./client_pda data/pda1_spec.txt < data/pda1_inp01.txt
#
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/* YOUR CODE GOES HERE */

func main() {
	initLogging()

	var buf []byte
	var msg string
	var err error

	cliArgs := getCliArgs()
	fname := cliArgs[0]
	buf, err = ioutil.ReadFile(fname)
	check(err)
	msg = readAllStdin()

	fmt.Printf("CLIargs:: %v\nSPEC:: %s\nINPUT:: %s\n", cliArgs, string(buf), msg)

	xpda := PDA_x{}
	mystring := string(buf)
	xpda.open(mystring)
	CLog.Printf("Source %s\n", xpda.source())
	fmt.Println("Valid PDA?", xpda.isValid())

	// check the PDA on the sample input
	tokens := strings.Fields(msg)
	fmt.Printf("token stream %v\n", tokens)
	/*
		xpda.feed("0")
		xpda.feed("0")
		xpda.feed("1")
		xpda.feed("1")
	*/
	/* YOUR CODE GOES HERE */

	for i := range tokens {
		if i != len(tokens)-1 && !xpda.isHang() {
			a := xpda.feed(tokens[i])
			_ = a
		} else {
			xpda.Eoi = true
			a := xpda.feed(tokens[i])
			_ = a
		}
	}

	fmt.Println("declaring end-of-input")
	xpda.noMore()
	m := 3
	fmt.Println("###### Standard Error ################")
	fmt.Fprintf(os.Stderr, "\nView clock=%d currentState=%s peek(%1d)=%v stack=%v\n",
		xpda.clock(), xpda.control(), m, xpda.peek(m), xpda.peek(-1)) //printed as Std Error

	fmt.Println("\n###### Standard Output ################")
	fmt.Fprintf(os.Stdout, "PDA Name=%v Accepted=%v clock=%d\n", xpda.code.Name, xpda.isAccepted(), xpda.clock()) //Printed as Std Out
	xpda.close()
	fmt.Println("End of Program")
}
