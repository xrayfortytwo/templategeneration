# Non of this is nearly ready, just to learn new cool things.

## Motivation

## Architecture and dataflow

![dataflow](https://raw.githubusercontent.com/xrayfortytwo/templategeneration/master/doc/flowdoku001.png)

## Usage
### Install with go on your system
    go get github.com/xrayfortytwo/templategeneration

### Execution
    go to "/templategeneration" directory
    
    go run main.go (-filename=<path to input *.g4 file> | -destination="<path to store output file>")*
    
### Flags
    -filename=<path to input *.g4 file>   
    -destination=<path to output dir>
    DEFALUT:
        filename=./input/fsml.g4
        destination=./tmp/template.temp

### TODOs and flaws

    - Only object grammars are excepted (no "|" in grammar rule)
    - For a good result in json representation elements in rules need to be distinct. 
    Unless a bit more manual enhancement of the grammar is needed.

### Related work

    - "A Formal Way from Text to Code Templates" by Guido Wachsmuth (Humboldt-Universität zu Berlin)
