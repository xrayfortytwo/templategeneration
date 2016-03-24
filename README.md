# Non of this is nearly ready, just a to learn new cool things.

## Motivation

## Architecture and dataflow

![dataflow](https://github.com/xrayfortytwo/templategeneration/doc/flowdoku001.png)

## Usage
### Install with go on your system
    `go get github.com/xrayfortytwo/templategeneration`

### Execution
    go to `/templategeneration` directory
    
    `go run main.go (-filename=<path to input *.g4 file> | -destination="<path to store output file>")*`
    
### Flags
    `-filename=<path to input *.g4 file>`   
    `-destination=<path to output dir>`
    DEFALUT:
        filename=./input/fsml.g4
        destination=./tmp/template.temp

### TODOs and flaws
    * Only Object Grammars are excepted (no `|` in grammar rule)
    * for a good result in json representation elements in rules need to be distinct. 
      Unless a bit more manual enhancement of the grammar is needed.

### Related work
    * "A Formal Way from Text to Code Templates" by Guido Wachsmuth (Humboldt-Universität zu Berlin)
