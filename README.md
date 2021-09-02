# Push-Down-Automata-with-Google-Go

This distributed automata system processes the input string and to verifies it against a PushDown Automata.\
The automata program requires the following:

1. A JSON file which defines the grammar for the push down automata\
2. An optional text file consisting of the string to be processed. User can provide input with standard input interface

The output of the program states if the provided string was accepted or rejected by the Push Down Automata. System also prints all the intermediate stages which were taken while processing each input token.

**How to run the program?**\
Automata System is implemented using a Golang program which can be run using following command in command prompt.
Please make sure that we open command prompt in the folder of program.
``` javascript
go run client_pda.go pda.go collection_functions.go utils.go -- data/pda1_spec.txt < data/pda1_inp01.txt
```
Here we have data folder which contains json specification of PDA. 
As sample we have also added two PDA’s .json 

Along with all the .go files we have added two bash scripts.
These scripts run the same command on the bash.

We can observe bash output in the prompt.
