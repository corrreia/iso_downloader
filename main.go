package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/SonarBeserk/gousbdrivedetector"
	"github.com/cavaliergopher/grab/v3"
	"github.com/corrreia/distro_downloader/download"
)

type Config struct {
	//array of Type
	Type []Type `json:"types"`
}

type Type struct {
	TypeName string `json:"type_name"`
	//array of OS
	OS []OS `json:"os"`
}

type OS struct {
	OSName string `json:"os_name"`
	DownloadURL string `json:"download_url"`
}

const configJSON = "config.json"

func main() {
	t := flag.Int("t", 5, "number of threads to download ISOs")
	flag.Parse()

	fmt.Println("Welcome to the ISO downloader!")

	Download := download.NewDownload()
	arrayOfDrives, err := getDrives()
	if err != nil {
		fmt.Println(err)
		return
	}

	arrayOfIsos := make([][]string, 0)

	fmt.Println("Getting ISOs...")
	dlType := reflect.TypeOf(Download)
	for i := 0; i < dlType.NumMethod(); i++ {

    	method := dlType.Method(i)
		ret := method.Func.Call([]reflect.Value{reflect.ValueOf(Download)})[0].Interface()

		if len(ret.([]string)) == 0 {
			fmt.Println("Something went wrong with " + method.Name)
			continue
		}

		arrayOfIsos = append(arrayOfIsos, ret.([]string))
	}

	fmt.Println("Available ISOs:\n")

	//display all the ISOs
	for i := 0; i < len(arrayOfIsos); i++ {
		fmt.Println(i+1, ":" , arrayOfIsos[i][3], arrayOfIsos[i][1], arrayOfIsos[i][2])
	}

	//ask user which ISOs to download (eg. 0 1 2 3, or all)
	fmt.Println("Which ISOs would you like to download? (eg. 0 1 2 3 / all)")

	var input string
	fmt.Scanln(&input)

	chosenIsos := make([]int, 0)

	//download the ISOs
	if input == "all" {
		for i := 0; i < len(arrayOfIsos); i++ {
			chosenIsos = append(chosenIsos, i)
		}

	} else {
		chosenIndexes := strings.Split(input, " ")
		for i := 0; i < len(chosenIndexes); i++ {
			index, err := strconv.Atoi(chosenIndexes[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			chosenIsos = append(chosenIsos, index-1)
		}
	}

	//chose a drive letter
	fmt.Println("Available drives:")

	for i := 0; i < len(arrayOfDrives); i++ {
		fmt.Println(arrayOfDrives[i])
	}

	fmt.Println("Which drive would you like to download the ISOs to?")
	var drive string
	fmt.Scanln(&drive)

	//begin downloading in t number threads
	// only use max t threads to download
	// path := drive + arrayOfIsos[chosenIsos[j]][3] + "/" + arrayOfIsos[chosenIsos[j]][1] + arrayOfIsos[chosenIsos[j]][2] + ".iso"
	// url := arrayOfIsos[chosenIsos[j]][0]

	fmt.Println("Thread count:", *t)

	for i := 0; i < len(chosenIsos); i++ {
		go downloadDistro(drive + arrayOfIsos[chosenIsos[i]][3] + "/" + arrayOfIsos[chosenIsos[i]][1] + arrayOfIsos[chosenIsos[i]][2] + ".iso", arrayOfIsos[chosenIsos[i]][0])
	}

	for{}
}

func downloadDistro(path string, url string) error {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(path, url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	Loop: for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		return err
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	return nil
}

func getDrives() ([]string, error) {
	return usbdrivedetector.Detect()
}