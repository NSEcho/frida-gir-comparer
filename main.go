package main

import (
	"fmt"
	"github.com/nsecho/fgcomparer/comparer"
	"github.com/nsecho/fgcomparer/helper"
	"github.com/nsecho/fgcomparer/parser"
	"os"
	"path/filepath"
	"sync"
)

const (
	kitName = "frida-core-%s-macos-arm64"
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
	downloadKits(oldVersion, newVersion, outdir, &wg)

	oldKitName := fmt.Sprintf(kitName, oldVersion)
	newKitName := fmt.Sprintf(kitName, newVersion)

	oldGir := filepath.Join(outdir, oldKitName, "frida-core.gir")
	newGir := filepath.Join(outdir, newKitName, "frida-core.gir")

	oldP, err := parser.NewParser(oldGir)
	if err != nil {
		panic(err)
	}

	newP, err := parser.NewParser(newGir)
	if err != nil {
		panic(err)
	}

	c := comparer.NewComparer(oldP, newP)
	c.Compare()
	fmt.Println(c.String())
}

func downloadKits(oldVersion, newVersion, outdir string, wg *sync.WaitGroup) {
	oldKitName := filepath.Join(outdir, fmt.Sprintf(kitName, oldVersion), "libfrida-core.a")
	newKitName := filepath.Join(outdir, fmt.Sprintf(kitName, newVersion), "libfrida-core.a")

	if _, err := os.Stat(oldKitName); os.IsNotExist(err) {
		wg.Add(1)
		go download(oldVersion, outdir, wg)
	} else {
		fmt.Printf("[*] Found devkit for %s version\n", oldVersion)
	}

	if _, err := os.Stat(newKitName); os.IsNotExist(err) {
		wg.Add(1)
		go download(newVersion, outdir, wg)
	} else {
		fmt.Printf("[*] Found devkit for %s version\n", newVersion)
	}

	wg.Wait()
}

func download(version, outdir string, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	for dl := range helper.DownloadAll(version, outdir) {
		if dl != "" {
			total += 1
		}
		fmt.Printf("[*] %s finished download; parsing gir file\n", dl)
	}
	if total != helper.Count() {
		fmt.Fprintf(os.Stderr, "[-] Did not download all kits for %s\n", version)
	}
}
