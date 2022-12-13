package download

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/k3a/html2text"
)

func (d Download) GetDebianLatest() ([]string, error) {
	versionLink := "https://www.debian.org/CD/"

	http, err := http.Get(versionLink)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	split := strings.Split(text, "Latest official release of the \"stable\" CD/DVD images: ")

	version := strings.Split(split[1], " ")[0]

	version = version[:len(version)-1]

	//use https://cdimage.debian.org/debian-cd/current-live/amd64/iso-hybrid/debian-live-11.5.0-amd64-cinnamon.iso
	link := fmt.Sprintf("https://cdimage.debian.org/debian-cd/current-live/amd64/iso-hybrid/debian-live-%s-amd64-cinnamon.iso", version)
	return []string{link, "Debian", version, "Linux"}, nil
}