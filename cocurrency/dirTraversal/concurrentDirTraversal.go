package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkDir1(dir string, fileSizes chan<- int64) {
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir1(subDir, fileSizes)
		} else {
			f, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot get file Info: %v", err)
			}
			fileSizes <- f.Size()
		}
	}
}

func dirEntries1(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	return entries
}

func main1() {
	flag.Parse() 
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	go func() {
		for _, d := range roots {
			walkDir1(d, fileSize)
		}
		close(fileSize)
	}()

	var numFiles, numBytes int64 
	for size := range fileSize {
		numFiles++
		numBytes += size
	}
	printDiskUsage1(numFiles , numBytes)
}

func printDiskUsage1(nfile, nBytes int64) {
	fmt.Printf("%d files %.1f \n", nfile, float64(nBytes))
}


