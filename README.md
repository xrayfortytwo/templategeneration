# Non of this is nearly ready, just to learn new cool things.

## Motivation
   
The main idea of this application is to extract all data needed to create a code template from its Antlr4 grammar.
To achieve this the grammar is scanned and divided into different types of rules.

These classes of rules are nonterminal rules, kleene rules, morphen rules and lexer rules.
Nonterminal rules are those who include nonterminals except kleene rules. However, kleene rules are rules who 
contain tokens with a kleene star "*" applied, thus a list is needed to fill the template.
Morphen rules are nonterminal rules without a defining rule. Finally there are some lexer rules who specify 
tokens in an antrl4 grammar.

By extracting these rules it is possible to create a template for the language defined by the grammar.
The meta-template, to create the output-template, consists of small templates for the different kinds of Tokens 
(e.g 'keywords', '+', token), by executing the meta-templates a small part of the template is written into a buffer. 
So ultimately a template is created.

## Architecture and dataflow

Simple overview:
![dataflow](https://raw.githubusercontent.com/xrayfortytwo/templategeneration/master/doc/flowdoku001.png)

More detailed overview:
![dataflow](https://raw.githubusercontent.com/xrayfortytwo/templategeneration/master/doc/flowdoku002.png)

## Usage
### Install with go on your system
    go get github.com/xrayfortytwo/templategeneration

### Makefile
This will generate the [template](https://github.com/xrayfortytwo/templategeneration/blob/master/tmp/temp.template) and the [POJO](https://github.com/xrayfortytwo/templategeneration/tree/master/tmp) for [fsml.g4](https://github.com/xrayfortytwo/templategeneration/blob/master/input/fsml.g4)

    cd to "/templategeneration" dir.
    make test_run

### Go Execution
    cd to "/templategeneration" directory
    go run main.go (-filename=<path to input *.g4 file> | -destination="<path to store output file>")*
    
### Flags
    -filename=<path to input *.g4 file>   
    -destination=<path to output dir>
    DEFALUT:
        filename=./input/fsml.g4
        destination=./tmp/template.temp

### TODOs and flaws

    - Only object grammars are excepted (no "|" in grammar rule)

### Related work

    - "A Formal Way from Text to Code Templates" by Guido Wachsmuth (Humboldt-Universität zu Berlin)
