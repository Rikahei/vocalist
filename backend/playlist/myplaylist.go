package playlist

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type List2 struct {
	VideoId   string `json:"videoId"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

func GetMyList(getJson echo.Context) (err error) {
	u := &[]List2{}
	if err = getJson.Bind(u); err != nil {
		return
	}
	fmt.Printf("%v", u)
	return getJson.JSON(http.StatusOK, u)
}
