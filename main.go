package main

/*
   Main entry point of the application, with flags
*/

import (
	"flag"
	"github.com/templategeneration/adaption"
	"github.com/templategeneration/genjson"
	"github.com/templategeneration/utils"
)

const (
	LOGPATH              = "./logs/log"
	DEFAULT_FILE         = "./input/fsml.g4"
	DESTINATION_TEMPLATE = "./tmp/temp.template"
	DESTINATION_JSON     = "./tmp/jsondata.json"
)

func main() {
	var fileName = flag.String("filename", DEFAULT_FILE, "specfiy file location")
	var destitemp = flag.String("destination_temp", DESTINATION_TEMPLATE, "specify temp. output")
	var destijson = flag.String("destination_json", DESTINATION_JSON, "specify json output")
	util.CleanUp(*destitemp, *destijson)
	flag.Parse()
	util.NewLog(LOGPATH)
	util.Log.Println("INFO: Read: ", *fileName)
	input := util.IOReadFile(*fileName)
	var s = genjson.GenJson(destijson)
	var t = &adaption.GenTemp{s}
	adaption.RunAdaption(&input, destitemp, t)
	s.Decode()
}
