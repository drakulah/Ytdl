package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"ytdl/conv"
	"ytdl/log"
	"ytdl/prog"
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
		log.E("invalid format provided")
	}

	fmt.Println("[Fetch] Preparing download request")

	client := &http.Client{}
	req, createReqErr := http.NewRequest("GET", ytInfo.AudioURL, nil)

	if createReqErr != nil {
		log.E("could not create a request")
		return
	}

	var file *os.File
	var offset int64 = 0
	var offsetEnd int64 = 1
	var highWaterMark int64 = 496 * 1000

	prog := prog.New()

	for {
		if offset > offsetEnd {
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-", offset))
		} else {
			req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", offset, offsetEnd))
		}
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			log.E("could npt make a request")
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
			log.E("error writing stream")
			return
		}

		offset += resp.ContentLength
		offsetEnd = offset + highWaterMark

		rng := resp.Header.Get("Content-Range")
		size := conv.StringToInt64(strings.Split(rng, "/")[1])

		if offsetEnd > size {
			offsetEnd = size
		}

		prog.SetCurrent(offset)
		prog.SetTotal(size)
		prog.Display()
	}

	fmt.Println("\n[Done] Downloaded successfully")
}

func ToValidFilename(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	filename := reg.ReplaceAllString(s, "_")
	filename = strings.Trim(filename, "_")

	return filename
}
