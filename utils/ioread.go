package util

import (
	"io/ioutil"
	"log"
	"os"
)

type OutputUtil struct {
	DestinationTemplate string
	DestinationJson     string
}

func (o OutputUtil) Write(b []byte) (int, error) {
	f, err := os.OpenFile(o.DestinationTemplate, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}
	f.Close()
	return -1, nil
}

func CleanUp(args ...string) {
	for _, s := range args {
		if _, err := os.Stat(s); err == nil {
			check(os.Remove(s))
			log.Println("Deleted", s)
		}
		log.Println("File created: ", s)
		os.Create(s)
	}
}

func IOReadFile(filepath string) []byte {
	input, err := ioutil.ReadFile(filepath)
	check(err)
	return input
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
