package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	inputBinaryFile  = flag.String("input-file", "", "input bin file")
	xorKey           = flag.Int("xor-key", 8, "input xor key")
	outputBinaryFile = flag.String("output-file", "output.bin", "input bin file")
	appendFlag = flag.String("append-flag", "A", "flag for file head, 1 bytes")
)

func usage() {
	fmt.Println("[gxor] Xor Binary file")
	fmt.Println("Usage : gxor -input-file payload.bin -output-file out.bin -xor-key 10 -appendFlag A")
}

func main() {
	flag.Parse()
	if *inputBinaryFile == "" {
		usage()
		return
	}
	fp, err := os.Open(*inputBinaryFile)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	fileInfo, err := fp.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Size:", fileInfo.Size())
	fileSize, _ := strconv.Atoi(strconv.FormatInt(fileInfo.Size(), 10))
	data := make([]byte, fileInfo.Size())
	fp.Read(data)
	for i := 0; i < fileSize; i++ {
		data[i] ^= byte(*xorKey)
	}
	newFile, err := os.Create(*outputBinaryFile)
	defer newFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	newFile.Write([]byte(*appendFlag))
	newFile.Write(data)
	fmt.Println("[*]Flag Size:", len(*appendFlag))
	fmt.Println("[*]Output:", *outputBinaryFile)
}
