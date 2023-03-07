package ytmusic

import "fmt"

func ContextAndroidYouTubeMusic(videoId string) string {
	return fmt.Sprintf(`{
		"videoId": "%s",
		"context": {
			"client": {
				"androidSdkVersion": 31,
				"gl"               : "US",
				"hl"               : "en",
				"clientVersion"    : "5.26.1",
				"clientName"       : "ANDROID_MUSIC",
				"visitorData"      : "%s"
			}
		}
	}`, videoId, visitorData)
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
