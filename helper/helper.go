package helper

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/ulikunitz/xz"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	releaseURL = "https://github.com/frida/frida/releases/download/%s/frida-core-devkit-%s-%s-%s.tar.xz"
)

var kit = "macos"
var arch = "arm64"

func Count() int {
	return 1
}

// DownloadAll downloads all the releases defined
func DownloadAll(version, outdir string) <-chan string {
	ch := make(chan string)
	go startDownload(version, outdir, ch)
	return ch
}

func startDownload(version, outdir string, ch chan<- string) {
	var wg sync.WaitGroup
	wg.Add(Count())
	go download(version, kit, arch, outdir, &wg, ch)
	wg.Wait()
	close(ch)
}

// Download downloads single core devkit for host and arch
func Download(version, kit, arch, outdir string) error {
	ch := make(chan string)
	go func() {
		<-ch
	}()
	return downloadItem(version, kit, arch, outdir, ch)
}

func download(version, kit, arch, outdir string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	downloadItem(version, kit, arch, outdir, ch)
}

func downloadItem(version, host, arch, outdir string, ch chan<- string) error {
	corePath := fmt.Sprintf("frida-core-%s-%s-%s", version, host, arch)
	dp := filepath.Join(outdir, corePath)
	fmt.Printf("[*] Downloading: frida-core %s for %s %s to %s/\n",
		version, host, arch, dp)
	frida := fmt.Sprintf(releaseURL, version, version, host, arch)
	resp, err := http.Get(frida)
	if err != nil {
		return err
	}

	r, err := xz.NewReader(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		resp.Body.Close()
		ch <- ""
		return err
	}

	resp.Body.Close()

	tr := tar.NewReader(buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			ch <- ""
			return err
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			dirPath := filepath.Join(outdir, corePath, hdr.Name)
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				ch <- ""
				return err
			}
		case tar.TypeReg:
			filePath := filepath.Join(outdir, corePath, hdr.Name)
			newFile, err := os.Create(filePath)
			if err != nil {
				ch <- ""
				return err
			}
			if _, err := io.Copy(newFile, tr); err != nil {
				newFile.Close()
				ch <- ""
				return err
			}
			newFile.Close()
		}
	}

	finished := fmt.Sprintf("frida-core-%s-%s-%s", version, host, arch)
	ch <- finished
	return nil
}
