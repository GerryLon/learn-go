package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 读取文件的各种方式
//
// 运行方式: go build main.go; ./main.exe

// * 整个文件读到内存，适用于文件较小的情况
func readAllIntoMemory(filename string) (content []byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

// * 一块一块地读取, 即给一个缓冲, 分多次读到缓冲中
func readByBlock(filename string) (content []byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	const bufferSize = 64 // 缓冲大小, 每次读取64个字节

	buffer := make([]byte, bufferSize)
	for {
		// 注意这里要取bytesRead, 否则有问题
		bytesRead, err := fp.Read(buffer) // 文件内容读取到buffer中
		content = append(content, buffer[:bytesRead]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return nil, err
			}
		}
	}
	return
}

// 逐行读取, 一行是一个[]byte, 多行就是[][]byte
func readByLine(filename string) (lines [][]byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)

	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			lines = append(lines, line)
		}
	}

	return
}

func main() {
	const testFileName = "ls-al.txt"
	// fmt.Println(file.GetCurrentDirectory())

	// *
	content, err := readAllIntoMemory(testFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", content)

	// *
	content, err = readByBlock(testFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", content)

	// *
	lines, err := readByLine(testFileName)
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range lines {
		fmt.Printf("%d %s\n", i+1, line)
	}
}
