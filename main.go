package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := os.Args[1]
	stem, _, found := strings.Cut(filepath, ".asm")
	if !found {
		log.Fatal("file must be a .asm")
	}
	stem += ".hack"

	outFile, err := os.Create(stem)
	if err != nil {
		log.Fatal("Error creating output file")
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	inFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	p := newParser(inFile)

	for p.hasMoreLines() {
		p.advance()

		out := ""
		switch p.instructionType() {
		case cInstruction:
			out = "111"
			out += cMap[p.comp]
			out += dMap[p.dest]
			out += jMap[p.jump]
		case aInstruction:
			addr, err := strconv.Atoi(p.symbol())
			if err != nil {
				log.Fatal(err)
			} else {
				out = strconv.FormatInt(int64(addr), 2)
				for len(out) < 16 {
					out = "0" + out
				}
			}
		case lInstruction:
			// Do nothing
		}
		_, err := fmt.Fprintln(w, out)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
}
