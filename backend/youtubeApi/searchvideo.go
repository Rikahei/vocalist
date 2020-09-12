package youtubeApi

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type VideoList struct {
	VideoId   string `json:"videoId"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

var (
	// query = flag.String("query", "初音ミク", "Search term")
	maxResults = flag.Int64("max-results", 9, "Max YouTube results")
)

const developerKey = "AIzaSyAvUAHwvmcvxgEQwi75SR30Om3o1ljF5vU"

func searchVideo(input string) map[string]string {

	queryCmd := flag.NewFlagSet("sKeyword", flag.ExitOnError)
	query := queryCmd.String("query", input, "Search term")

	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		fmt.Println(err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	// channels := make(map[string]string)
	// playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			details := VideoList{
				VideoId:   item.Id.VideoId,
				Title:     item.Snippet.Title,
				Thumbnail: item.Snippet.Thumbnails.High.Url,
			}
			jsonList, _ := json.Marshal(details)
			videos[item.Id.VideoId] = string(jsonList)
			// case "youtube#channel":
			// 	channels[item.Id.ChannelId] = item.Snippet.Title
			// case "youtube#playlist":
			// 	playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	// printIDs("Channels", channels)
	// printIDs("Playlists", playlists)

	return videos
}

func headers(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/"):]
	fmt.Println(title)
	searchVideo(title)
}

func GetKeyword(keyword string) map[string]string {

	line := keyword
	result := searchVideo(line)
	// printIDs("Videos", result)
	return result

}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
