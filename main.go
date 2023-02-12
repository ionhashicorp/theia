package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// func UnGzip(source, target string) error {
// 	reader, err := os.Open(source)
// 	if err != nil {
// 		return err
// 	}
// 	defer reader.Close()

// 	archive, err := gzip.NewReader(reader)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(archive.Name)
// 	defer archive.Close()

// 	target = filepath.Join(target, archive.Name)
// 	writer, err := os.Create(target)
// 	if err != nil {
// 		return err
// 	}
// 	defer writer.Close()

// 	_, err = io.Copy(writer, archive)
// 	return err
// }

func main() {
	// sourceAbsFilePath, err := filepath.Abs(os.Args[0])
	sourceAbsFilePath, err := filepath.Abs("../consul-debug-2023-01-26T18-00-07+0100.tar.gz")
	if err != nil {
		fmt.Println(err)
	}
	// reader ->> open file
	reader, err := os.Open(sourceAbsFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	// pass file to archive
	archive, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	defer archive.Close()

	target := filepath.Join("/Users/ion/Documents/golang/", "consul-debug-2023-01-26T18-00-07+0100.tar")
	writer, err := os.Create(target)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		fmt.Println(err)
	}
	// UNTAR

	reader, err = os.Open(target)
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		path := filepath.Join("/Users/ion/Documents/golang/", header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				fmt.Println(err)
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			fmt.Println(err)
		}
	}
}
