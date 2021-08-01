package main

import (
	"bufio"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func roomsInit() ([]*Room, error) {
	r := []*Room{}
	roomdoc, err := readLines("/usr/go/src/gopherit/GopherIT/roomdoc.txt")
	if err != nil {
		return r, err
	}
	i := 0
	index := 0
	for _, str := range roomdoc {
		if i == 0 {
			r = append(r, &Room{})
			r[index].roomInitId(str)
			i++
		} else {
			if i == 1 {
				r[index].roomInitDesc(str)
				i++
			} else {
				if i == 2 {
					links := strings.Split(str, ",")
					for _, link := range links {
						vri := strings.Split(link, ":")
						v, ri := vri[0], vri[1]
						r[index].addLink(v, ri)
					}
					index++
					i = 0
				}
			}
		}

	}
	return r, err
}
