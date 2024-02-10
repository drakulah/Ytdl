package ytmusic

import "fmt"

func ContextAndroidYouTubeMusic(videoId string) map[string]interface{} {
	return map[string]interface{}{
		"videoId": videoId,
		"context": map[string]interface{}{
			"client": map[string]interface{}{
				"androidSdkVersion": 32,
				"gl":                "US",
				"hl":                "en",
				"clientVersion":     "6.39.52",
				"clientName":        "ANDROID_MUSIC",
				"visitorData":       visitorData,
			},
		},
	}
}

func ContextAndroidYouTube(videoId string) string {
	return fmt.Sprintf(`{
		"videoId": "%s",
		"context": {
			"client": {
				"androidSdkVersion": 31,
				"gl"               : "US",
				"hl"               : "en",
				"clientVersion"    : "17.36.4",
				"clientName"       : "ANDROID",
				"visitorData"      : "%s"
			}
		}
	}`, videoId, visitorData)
}

func ContextWebYouTube(videoId string) string {
	return fmt.Sprintf(`{
		"videoId": "%s",
		"context": {
			"client": {
				"gl"               : "US",
				"hl"               : "en",
				"clientVersion"    : "2.20220918",
				"clientName"       : "WEB",
				"visitorData"      : "%s"
			}
		}
	}`, videoId, visitorData)
}
