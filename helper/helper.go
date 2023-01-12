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
	"time"
)

const (
	releaseURL = "https://github.com/frida/frida/releases/download/%s/frida-core-devkit-%s-%s-%s.tar.xz"
)

var releases = map[string][]string{
	"macos": []string{"arm64", "arm64e", "x86_64"},
}

func DownloadAll(version string) {
	client := newClient(60)

	count := 0
	for _, arches := range releases {
		count += len(arches)
	}

	var wg sync.WaitGroup
	wg.Add(count)

	for kit, arches := range releases {
		for _, arch := range arches {
			go download(version, kit, arch, client, &wg)
		}
	}

	wg.Wait()
}

func Download(version, kit, arch string) error {
	client := newClient(60)
	return downloadItem(version, kit, arch, client)
}

func download(version, kit, arch string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	client = newClient(60)
	if err := downloadItem(version, kit, arch, client); err != nil {
		panic(err)
	}
}

func downloadItem(version, host, arch string, client *http.Client) error {
	fmt.Printf("[*] Downloading: frida-core %s for %s %s\n",
		version, host, arch)
	frida := fmt.Sprintf(releaseURL, version, version, host, arch)
	resp, err := client.Get(frida)
	if err != nil {
		return err
	}

	corePath := fmt.Sprintf("frida-core-%s-%s", host, arch)
	fmt.Printf("[*] Destination dir: %s\n", corePath)

	r, err := xz.NewReader(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		resp.Body.Close()
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
			return err
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			dirPath := filepath.Join(corePath, hdr.Name)
			if err := os.MkdirAll(dirPath, 0777); err != nil {
				return err
			}
		case tar.TypeReg:
			filePath := filepath.Join(corePath, hdr.Name)
			newFile, err := os.Create(filePath)
			if err != nil {
				return err
			}
			if _, err := io.Copy(newFile, tr); err != nil {
				newFile.Close()
				return err
			}
			newFile.Close()
		}
	}

	return nil
}

func newClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}
