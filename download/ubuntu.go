package download

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/k3a/html2text"
)

func (d Download) GetUbuntuLTS() ([]string, error) {
	url := "https://ubuntu.com/download/desktop"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	re := regexp.MustCompile(`(?m)Ubuntu (([0-9])+.?)+(?P<name> LTS)`)

	match := re.FindStringSubmatch(text)

	version := strings.Split(match[0], " ")[1]
	
	//link is https://releases.ubuntu.com/%s/ubuntu-%s-desktop-amd64.iso
	link := fmt.Sprintf("https://releases.ubuntu.com/%s/ubuntu-%s-desktop-amd64.iso", version, version)
	return []string{link, "Ubuntu_LTS", version, "Linux"}, nil
}

func (d Download) GetUbuntuLatest() ([]string , error) {
	url := "https://ubuntu.com/download/desktop"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	re := regexp.MustCompile(`(?m)Ubuntu (([0-9])+.?)+`)
	reLTS := regexp.MustCompile(`(?m)Ubuntu (([0-9])+.?)+(?P<name> LTS)`)

	versionLTS := strings.Split(reLTS.FindStringSubmatch(text)[0], " ")[1]
	version := "00.00"

	for _, match := range re.FindAllStringSubmatch(text, -1) {
		versionT := strings.Split(match[0], " ")[1]
		if versionT > versionLTS {
			version = versionT
		}
	}

	link := fmt.Sprintf("https://releases.ubuntu.com/%s/ubuntu-%s-desktop-amd64.iso", version, version)
	return []string{link, "Ubuntu", version, "Linux"}, nil
}

func (d Download) GetUbuntuServerLTS() ([]string, error) {
	url := "https://ubuntu.com/download/server"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	re := regexp.MustCompile(`(?m)Ubuntu Server (([0-9])+.?)+(?P<name> LTS)`)

	match := re.FindStringSubmatch(text)

	version := strings.Split(match[0], " ")[2]
	
	//link is https://releases.ubuntu.com/%s/ubuntu-%s-live-server-amd64.iso
	link := fmt.Sprintf("https://releases.ubuntu.com/%s/ubuntu-%s-live-server-amd64.iso", version, version)
	return []string{link, "Ubuntu_Server_LTS", version, "Linux"}, nil
}

func (d Download) GetUbuntuServerLatest() ([]string, error) {
	url := "https://ubuntu.com/download/server"

	http, err := http.Get(url)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := html2text.HTML2Text(string(body))

	re := regexp.MustCompile(`(?m)Ubuntu Server (([0-9])+.?)+`)
	reLTS := regexp.MustCompile(`(?m)Ubuntu Server (([0-9])+.?)+(?P<name> LTS)`)

	versionLTS := strings.Split(reLTS.FindStringSubmatch(text)[0], " ")[2]
	version := "00.00"

	for _, match := range re.FindAllStringSubmatch(text, -1) {
		versionT := strings.Split(match[0], " ")[2]
		if versionT > versionLTS {
			version = versionT
		}
	}

	link := fmt.Sprintf("https://releases.ubuntu.com/%s/ubuntu-%s-live-server-amd64.iso", version, version)
	return []string{link, "Ubuntu_Server", version, "Linux"}, nil
}