package main

import (
	"backend/nicoApi"
	"net/http"

	"backend/playlist"
	"backend/youtubeApi"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func oauth2login(c echo.Context) error {
	loginlink := youtubeApi.GetUrl()
	println(loginlink)
	return c.JSON(http.StatusOK, loginlink)
}

func oauth2Token(c echo.Context) error {
	code := c.FormValue("code")
	println("code: ", code)

	// Exchange oauth code to youtube api
	mineResponse := youtubeApi.GetMineInfo(code)

	url := "http://localhost:8010/#/oauth2success?channelId=" + mineResponse.Items[0].Id +
		"&title=" + mineResponse.Items[0].Snippet.Title
	return c.Redirect(http.StatusMovedPermanently, url)
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func searchVideo(c echo.Context) error {
	keyword := c.FormValue("keyword")
	result := youtubeApi.GetKeyword(keyword)

	return c.JSON(http.StatusOK, result)
}

func nicoSearchVideo(c echo.Context) error {
	keyword := c.FormValue("keyword")
	result := nicoApi.NicoSearch(keyword)
	return c.JSON(http.StatusOK, result)
}

func updateMyList(c echo.Context) (err error) {
	result := playlist.GetMyList(c)
	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	// e.POST("/login", login)
	e.GET("/oauth2", oauth2login)
	e.GET("/oauth2token", oauth2Token)

	// Unauthenticated route
	e.GET("/", accessible)

	e.POST("/search", nicoSearchVideo)
	e.POST("/updateMyList", updateMyList)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
