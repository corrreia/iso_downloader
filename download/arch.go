package download

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/k3a/html2text"
)

func (d Download) GetArchLatest() ([]string, error) {
	versionLink := "https://www.archlinux.org/download/"

	http, err := http.Get(versionLink)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	split := strings.Split(text, "Current Release: ")

	version := strings.Split(split[1], " ")[0]

	//use https://geo.mirror.pkgbuild.com/iso/2022.12.01/archlinux-2022.12.01-x86_64.iso
	link := fmt.Sprintf("https://geo.mirror.pkgbuild.com/iso/%s/archlinux-%s-x86_64.iso", version, version)
	return []string{link, "Arch", version, "Linux"}, nil
}