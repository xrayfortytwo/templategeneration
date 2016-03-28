package util

import (
	"os"
)

type PrintJFile struct {
	FileName string
}

func (p PrintJFile) Write(b []byte) (int, error) {
	f, err := os.OpenFile("./tmp/"+p.FileName+".java", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	check(err)
	_, err = f.Write(b)
	check(err)
	f.Close()
	return -1, nil
}
