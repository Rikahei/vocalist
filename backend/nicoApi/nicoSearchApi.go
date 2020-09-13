package nicoApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type VideoList struct {
	VideoId   string `json:"videoId"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

func NicoSearch(keyword string) map[string]string {

	println(keyword)

	req, err := http.NewRequest("GET", "https://api.search.nicovideo.jp/api/v2/video/contents/search?q="+
		url.QueryEscape(keyword)+"&targets=title&fields=contentId,title,viewCounter,thumbnailUrl&filters[viewCounter][gte]=10000&_sort=-viewCounter&_offset=0&_limit=12&_context=apiguide", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// videos := make(map[string]string)

	body, err := ioutil.ReadAll(resp.Body)

	var decode struct {
		Data []struct {
			VideoId   string `json:"contentId"`
			Title     string `json:"title"`
			Thumbnail string `json:"thumbnailUrl"`
		}
	}

	err = json.Unmarshal(body, &decode)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(decode.Data[0].Title)

	videos := make(map[string]string)

	for _, item := range decode.Data {
		details := VideoList{
			VideoId:   item.VideoId,
			Title:     item.Title,
			Thumbnail: item.Thumbnail,
		}
		jsonList, _ := json.Marshal(details)
		videos[item.VideoId] = string(jsonList)
	}

	printIDs("Videos", videos)

	return videos
}

func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
