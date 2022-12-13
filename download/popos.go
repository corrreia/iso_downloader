package download

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func (d Download) GetPopOSLatest() ([]string, error) {
	versionLink := "https://blog.system76.com/"

	http, err := http.Get(versionLink)
	if err != nil { return nil, err }

	body, err := io.ReadAll(http.Body)
	if err != nil { return nil, err }

	text := string(body)

	re := regexp.MustCompile(`(?m)Pop!_OS (([0-9])+.?)+ LTS`)

	version := re.FindString(text)  //! very janky way to get the version, might be broken in the future

	version = strings.Split(version, " ")[1]

	//link https://iso.pop-os.org/22.04/amd64/intel/19/pop-os_22.04_amd64_intel_19.iso
	link := fmt.Sprintf("https://iso.pop-os.org/%s/amd64/intel/19/pop-os_%s_amd64_intel_19.iso", version, version)
	return []string{link, "Pop!_OS_LTS", version, "Linux"}, nil
}