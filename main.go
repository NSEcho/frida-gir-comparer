package main

import (
	"fmt"
	"github.com/lateralusd/fgcomparer/helper"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage:   %s oldVersion newVersion outdir/\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 16.0.7 16.0.8 outdir/\n", os.Args[0])
		os.Exit(1)
	}

	oldVersion := os.Args[1]
	newVersion := os.Args[2]
	outdir := os.Args[3]

	var wg sync.WaitGroup
	wg.Add(2)

	go download(oldVersion, outdir, &wg)
	go download(newVersion, outdir, &wg)

	wg.Wait()
}

func download(version, outdir string, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	for dl := range helper.DownloadAll(version, outdir) {
		if dl != "" {
			total += 1
		}
		fmt.Printf("%s finished download; parsing gir file\n", dl)
	}
	if total != helper.Count() {
		fmt.Fprintf(os.Stderr, "Did not download all kits for %s\n", version)
	}
}
