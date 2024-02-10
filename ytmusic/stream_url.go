package ytmusic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"ytdl/array"
	"ytdl/str"
)

type AdaptiveFormats struct {
	Itag             int    `json:"itag"`
	Url              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	AverageBitrate   int    `json:"averageBitrate"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength"`
	Quality          string `json:"quality"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
	InitRange        struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"initRange"`
	IndexRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"indexRange"`
}

type PlayerDetails struct {
	StreamingData struct {
		ExpiresInSeconds string            `json:"expiresInSeconds"`
		AdaptiveFormats  []AdaptiveFormats `json:"adaptiveFormats"`
	} `json:"streamingData"`
	VideoDetails struct {
		Title string `json:"title"`
	} `json:"videoDetails"`
}

type YouTubeInfo struct {
	Title    string
	AudioURL string
}

func GetAudioUrl(videoId string) (YouTubeInfo, error) {
	var ytInfo YouTubeInfo
	url := ""
	if len(videoId) == 0 {
		return ytInfo, errors.New("invalid video id")
	}

	body, err := json.Marshal(ContextAndroidYouTubeMusic(videoId))

	if err != nil {
		return ytInfo, err
	}

	reqUrl := "https://music.youtube.com/youtubei/v1/player?key=" + KEY_AND_MUSIC + "&prettyPrint=false"

	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(body))

	if err != nil {
		return ytInfo, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ytInfo, err
	}

	fmt.Println("[Fetch] Downloading API response")

	defer res.Body.Close()
	resBody, readErr := io.ReadAll(res.Body)
	var playerDetails PlayerDetails

	if readErr != nil {
		return ytInfo, readErr
	}

	fmt.Println("[Parse] Parsing response as json")

	if err := json.Unmarshal(resBody, &playerDetails); err != nil {
		return ytInfo, err
	}

	fmt.Println("[Fetch] Choosing best audio format")
	formats := playerDetails.StreamingData.AdaptiveFormats
	audioMimeType, _ := regexp.Compile(`^audio/(webm|mp4);`)

	formats = array.Filter(formats, func(e AdaptiveFormats) bool {
		return audioMimeType.MatchString(e.MimeType)
	})

	array.QuickSort(formats, func(a, b AdaptiveFormats) bool {
		return a.Bitrate > b.Bitrate
	})

	if len(formats) == 0 {
		return ytInfo, errors.New("formats not found")
	}

	url = strings.Trim(formats[0].Url, " ")

	if len(url) == 0 {
		return ytInfo, errors.New("got empty string as url")
	} else if !str.StartsWith(url, "https://") {
		return ytInfo, errors.New("got invalid stream url")
	}

	ytInfo.AudioURL = url
	ytInfo.Title = playerDetails.VideoDetails.Title

	return ytInfo, nil
}
