package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func isNull(val string) bool {
	if val == "" {
		return true
	}
	return false
}

// structure of a PDA specification
// each transition is a [5]string with 3 entries current (state, token, stacktop) and 2 entries next(state, stacktop)

type PDA struct {
	Name            string     `json:"name"`
	States          []string   `json:"states"`
	InputAlphabet   []string   `json:"inputAlphabet"`
	StackAlphabet   []string   `json:"stackAlphabet"`
	AcceptingStates []string   `json:"acceptingStates"`
	StartState      string     `json:"startState"`
	Transitions     [][]string `json:"transitions"`
	Eos             string     `json:"eos"`
}

func singlecontain(a []string, b string) bool { //To check whether starting state of PDA is valid
	var flag bool
	for i := range a {
		if a[i] == b {
			flag = true
			break
		}
	}
	if flag {
		return true
	} else {
		return false
	}
}

func multicontain(a []string, bc []string) bool { //To check whether accepting states are part of PDA states
	var flag []int
	for i := range bc {
		for j := range a {
			if bc[i] == a[j] {
				flag = append(flag, 1)
				break
			}
		}

	}
	for i := range flag {
		if flag[i] == 1 {
			continue
		} else {
			return false
		}
	}
	return true
}

func triplecontain(a []string, b [][]string) bool { //For validating Transitions with input and stack alphabets
	var flag []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < len(a); k++ {
				if b[i][j] == a[k] {
					flag = append(flag, 1)
					break
				}
			}
		}
	}
	for i := range flag {
		if flag[i] == 1 {
			continue
		} else {
			return false
		}
	}
	return true
}

func (p *PDA) isValid() bool { //Checks validity of Json specification of PDa with given conditions
	if singlecontain(p.States, p.StartState) && multicontain(p.States, p.AcceptingStates) && triplecontain(p.States, p.Transitions) || triplecontain(p.StackAlphabet, p.Transitions) || triplecontain(p.InputAlphabet, p.Transitions) {
		return true
	} else {
		return false
	}
}

// structure of configuration/state of a running PDA, i.e. "process"
type PDA_x struct {
	code    PDA      `json:"-"`       // ignore this field; internal/parsed PDA object
	Source  string   `json:"source"`  // the JSON string of the "program" of the PDA
	Eoi     bool     `json:"eoi"`     // notified of the end of the input stream?
	Control string   `json:"control"` // current state of control
	Stack   []string `json:"stack"`   // contents of the stack (top at the end)
	Clock   int      `json:"clock"`   // number of transitions taken
}

func (p *PDA_x) isHang() bool { //Checks whether PDA is hanged at non accepting states
	if (p.Control == "q2" && p.Eoi && p.Stack[len(p.Stack)-1] == "0") || (p.Control == "q3" && p.Stack[len(p.Stack)-1] == "$") {
		fmt.Println("PDA is Hanged at ->", p.Control)
		fmt.Println("PDA will be reset")
		fmt.Printf("\nPeek at Stack before the reset=%v \n", p.peek(-1))
		p.reset()
		p.close()
		return true
	} else {
		return false
	}
}

func (p *PDA_x) push(ele string) { //It pushes input token to stack top as input token is 0
	fmt.Println("Push -> ", ele)
	p.Stack = append(p.Stack, ele)
	fmt.Println("Pushed in Stack ->", p.Stack)
}

func (p *PDA_x) pop() { //It pops the Stack top element as 1 is encountered in input sequence
	p.Stack = p.Stack[:len(p.Stack)-1]
	fmt.Println("Popped ->", p.Stack)
}

func (p *PDA_x) marshal() []byte { //Marshals the json specification .ie. it encodes the json
	buf, err := json.Marshal(p)
	check(err)
	return buf
}

func (p *PDA_x) unmarshal(buf []byte) { //Unmarshals the Json specification of PDA it basically decodes the json
	err := json.Unmarshal([]byte(string(buf)), p)
	check(err)
}

func (p *PDA_x) isValid() bool { //Checks the validity of PDA with given conditions
	if singlecontain(p.code.States, p.code.StartState) && multicontain(p.code.States, p.code.AcceptingStates) && triplecontain(p.code.States, p.code.Transitions) || multicontain(p.code.States, p.code.InputAlphabet) || multicontain(p.code.States, p.code.StackAlphabet) {
		return true
	} else {
		return false
	}
}

func (p *PDA_x) open(spec string) { //Opens the json specification of pda, unmarshals it and sets up the PDA for operation
	err := json.Unmarshal([]byte(spec), &p.code)
	check(err)
	p.Stack = []string{}
	p.Clock = 0
	p.Control = "q1"
	p.Eoi = false
}

func (p *PDA_x) source() string { //Marshals json spcification of PDA AND PRINTS IT
	var SourceJson []byte
	SourceJson, err := json.Marshal(p.code)
	if err != nil {
		log.Println(err)
	}
	return string(SourceJson)
}

func (p *PDA_x) reset() { //Resets the PDA to default stae
	p.Control = p.code.StartState
	p.Eoi = false
	p.Stack = []string{}
	p.Clock = 0
}

func (p *PDA_x) clock() int { //Returns clock value of PDA
	return p.Clock
}

func (p *PDA_x) control() string { //Returns current state of PDA i.e. Control
	return p.Control
}

func (p *PDA_x) peek(k int) []string { //Peeks in to the stack depending upon the value passed to function
	if len(p.Stack) == 0 {
		var arr []string
		arr = append(arr, "Stack is Empty since matching number of zeros and ones in input or PDA is been Reset")
		return arr
	} else if k < 0 {
		return p.Stack
	} else {
		return p.Stack[0 : k+1]
	}
}

func (p *PDA_x) isAccepted() bool { //Function checks whether current input sequence is accepted by PDA or not
	var condition bool
	for i := range p.code.AcceptingStates {
		if strings.Contains(p.code.AcceptingStates[i], p.Control) && p.Control != "" {
			condition = true
			_ = condition
			break
		}
	}
	if condition {
		return true
	} else {
		return false
	}
}

func (p *PDA_x) noMore() int { //This function announces end of the input to PDA
	fmt.Println("END OF TOKEN STREAM")
	p.Eoi = true
	return p.clock()
}

func (p *PDA_x) feed(token string) int { // This function accepts a single token input from input files and depending upon state of PDA and input appropriate transition takes place.
	fmt.Println("Current Input Token to PDA-> ", token)
	if p.Control == "q1" { //Transition from q1 to q2 state
		p.Control = "q2"
		/*call push */
		p.push("$")
		p.push(token)
		fmt.Println("Transition to", p.Control)
		p.Clock = p.Clock + 1
		fmt.Printf("\nPeek at Stack after transition=%v \n", p.peek(-1))
		return 1
	}
	if token == "0" && p.Control == "q2" && p.Stack[len(p.Stack)-1] == "0" { //Transition from q2 to q2 state (Loop back to same state)
		p.Control = "q2"
		/*call push*/
		p.push(token)
		fmt.Println("Transition to", p.Control)
		fmt.Printf("\nPeek at Stack after transition=%v \n", p.peek(-1))
		return 0
	}
	if p.Control == "q2" && p.Eoi && p.Stack[len(p.Stack)-1] == "0" { //Checking whether hanged at q2 state
		p.isHang()
	}
	if token == "1" && p.Control == "q2" && p.Stack[len(p.Stack)-1] == "0" { //Transition from q2 to q3 state
		p.Control = "q3"
		p.Clock = p.Clock + 1
		/*call pop*/
		p.pop()
		fmt.Println("Transition to", p.Control)
		fmt.Printf("\nPeek at Stack after transition=%v \n", p.peek(-1))
		return 1
	}

	if token == "1" && p.Control == "q3" && p.Stack[len(p.Stack)-1] == "0" { //Transition from q3 to q3 state (Loop back to same state)
		p.Control = "q3"
		/*call pop*/
		p.pop()
		fmt.Println("Transition to", p.Control)
		fmt.Printf("\nPeek at Stack after transition=%v \n", p.peek(-1))
		if p.Eoi && p.Stack[len(p.Stack)-1] == "$" {
			p.Clock = p.Clock + 1
			p.Control = "q4"
			p.pop()
			fmt.Println("Transition to", p.Control)
			fmt.Printf("\nPeek at Stack after transition=%v \n", p.peek(-1))
			fmt.Println("Reached Accepted State Q4")
		}

	}
	if token == "1" && p.Control == "q3" && p.Stack[len(p.Stack)-1] == "$" { //Checking whether hanged at q3 state
		p.isHang()
	}

	return p.clock()
}

func (p *PDA_x) close() { //This closes PDA where all parameters are reset or set to default
	p.Eoi = false
	p.Control = ""
	p.Stack = []string{}
	p.Clock = 0
}
