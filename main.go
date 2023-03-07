package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"ytdl/log"
	"ytdl/ytmusic"
)

func main() {

	var videoId string
	var filterAV string
	var outputLoc string

	flag.StringVar(&videoId, "id", "", "YouTube's video id")
	flag.StringVar(&filterAV, "f", "a", "Audio or video to download")
	flag.StringVar(&outputLoc, "o", "./", "Location to save the downloads")

	flag.Parse()

	ytInfo, ytErr := ytmusic.GetAudioUrl(videoId)

	if ytErr != nil {
		log.E(ytErr.Error())
		return
	}

	filterAV = strings.ToLower(filterAV)

	if !regexp.MustCompile(`^(a|av|va|v)$`).MatchString(filterAV) {
		log.E("Invalid format provided")
	}

	client := &http.Client{}
	req, createReqErr := http.NewRequest("GET", ytInfo.AudioURL, nil)

	if createReqErr != nil {
		log.E("Couldn't create a request")
		return
	}

	var file *os.File
	var offset int64

	log.I("Starting download")

	for {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", offset))
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			log.E("Couldn't make a request")
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusPartialContent {
			break
		}

		if file == nil {
			file, err = os.Create(fmt.Sprintf("%s.opus", ToValidFilename(ytInfo.Title)))
			if err != nil {
				log.E(err.Error())
				return
			}

			defer file.Close()
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.E("Error writing stream")
			return
		}

		offset += resp.ContentLength
	}

	log.S("Successfully downloaded")
}

func ToValidFilename(name string) string {
	reg := regexp.MustCompile("^[a-zA-Z0-9 ]+")
	name = reg.ReplaceAllString(name, "")

	if len(name) > 255 {
		name = name[:255]
	}

	return name
}
