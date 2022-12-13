package download

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func (d Download) GetKaliLatest() ([]string, error) {
	versionLink := "https://www.kali.org/downloads/"

	http, err := http.Get(versionLink)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := string(body)

	re := regexp.MustCompile(`(?m)Kali Linux (([0-9])+.?)+`)

	version := re.FindString(text)
	
	version = strings.Split(version, " ")[2]

	//link https://cdimage.kali.org/kali-2022.4/kali-linux-2022.4-installer-netinst-amd64.iso
	link := fmt.Sprintf("https://cdimage.kali.org/kali-%s/kali-linux-%s-installer-netinst-amd64.iso", version, version)
	return []string{link, "Kali_Linux", version, "Linux"}, nil
}