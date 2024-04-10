package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func walkDir2(dir string, fileSizes chan<- int64) {
// 	for _, entry := range dirEntries(dir) {
// 		if entry.IsDir() {
// 			subDir := filepath.Join(dir, entry.Name())
// 			walkDir(subDir, fileSizes)
// 		} else {
// 			f, err := entry.Info()
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "cannot get file Info: %v", err)
// 			}
// 			fileSizes <- f.Size()
// 		}
// 	}
// }

// func dirEntries2(dir string) []os.DirEntry {
// 	entries, err := os.ReadDir(dir)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %v\n", err)
// 	}
// 	return entries
// }

// func main2() {
// 	verbose := flag.Bool("v", false, "show the verbose progress message")
// 	flag.Parse() 
// 	roots := flag.Args()
// 	fmt.Println(roots)
// 	// if len(roots) == 0 {
// 		roots = []string{"."}
// 	}

// 	fileSize := make(chan int64)
// 	go func() {
// 		for _, d := range roots {
// 			walkDir(d, fileSize)

// 		}
// 		close(fileSize)
// 	}()

// 	var tick <-chan time.Time
// 	if *verbose {
// 		tick = time.Tick(5000 * time.Millisecond)
// 	}
// 	var numFiles, numBytes int64 
// 	loop:
// 	for {
// 		select {
// 		case fs, ok := <- fileSize:
// 				if !ok {
// 					break loop
// 				}
// 				numFiles++
// 				numBytes += fs
// 			case <-tick:
// 				printDiskUsage(numFiles , numBytes)
// 		}
// 	}
// }

// func printDiskUsage(nfile, nBytes int64) {
// 	fmt.Printf("%d files %.1f \n", nfile, float64(nBytes))
// }


