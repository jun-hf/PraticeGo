package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subDir, fileSizes, wg)
		} else {
			f, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot get file Info: %v", err)
			}
			fileSizes <- f.Size()
		}
	}
}

func dirEntries(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	return entries
}

func main() {
	verbose := flag.Bool("v", false, "show the verbose progress message")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	wg := &sync.WaitGroup{}
	go func() {
		for _, d := range roots {
			wg.Add(1)
			go walkDir(d, fileSize, wg)
		}
	}()

	go func() {
		wg.Wait()
		close(fileSize)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(200 * time.Millisecond)
	}
	var numFiles, numBytes int64
	loop:
	for {
		select {
		case fs, ok := <- fileSize:
				if !ok {
					break loop
				}
				numFiles++
				numBytes += fs
			case <-tick:
				printDiskUsage(numFiles , numBytes)
		}
	}
	printDiskUsage(numFiles , numBytes)
}

func printDiskUsage(nfile, nBytes int64) {
	fmt.Printf("%d files %.1f \n", nfile, float64(nBytes))
}
