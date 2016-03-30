package main

/*
   Main entry point of the application, with flags
*/

import (
	"flag"
	"github.com/templategeneration/adaption"
	"github.com/templategeneration/genjson"
	"github.com/templategeneration/utils"
	"log"
	fpath "path/filepath"
)

const (
	DEFAULT_FILE         = "/input/fsml.g4"
	DESTINATION_TEMPLATE = "/tmp/temp.template"
	DESTINATION_JSON     = "/tmp/jsondata.json"
)

func main() {
	var cDir = detDir()
	var fileName = flag.String("filename", cDir+DEFAULT_FILE, "specfiy file location")
	var destitemp = flag.String("destination_temp", cDir+DESTINATION_TEMPLATE, "specify temp. output")
	var destijson = flag.String("destination_json", cDir+DESTINATION_JSON, "specify json output")
	util.CleanUp(*destitemp, *destijson)
	flag.Parse()
	log.Println("INFO: Read: ", *fileName)
	input := util.IOReadFile(*fileName)
	var s = genjson.GenJson(destijson)
	var t = &adaption.GenTemp{Gj: s}
	adaption.RunAdaption(&input, destitemp, t)
	s.Decode()
}

func detDir() string {
	dir, err := fpath.Abs("")
	check(err)
	log.Println(dir)
	return dir
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
