package download

import (
	"io"
	"net/http"
	"strings"

	"github.com/k3a/html2text"
)

// Download downloads the distro to the drive
func (d Download) GetFedoraWorkstationLatest() ([]string ,error) {
	url := "https://getfedora.org/en/workstation/download/"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	//from text, get the download link (its after "x86_64 Live ISO")
	link := strings.Split(strings.Split(text, "x86_64 Live ISO")[1], " ")[1]

	//get the version number
	version := strings.Split(link, "-")[4]

	return []string{link, "Fedora_Workstation", version, "Linux"}, nil
}

func (d Download) GetFedoraServerLatest() ([]string ,error) {
	url := "https://getfedora.org/en/server/download/"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	//from text, get the download link (its after "x86_64 Live ISO")
	link := strings.Split(strings.Split(text, "Standard ISO image for x86_64")[1], " ")[1]

	//get the version number
	version := strings.Split(link, "-")[4]

	return []string{link, "Fedora_Server", version, "Linux"}, nil
}