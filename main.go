package main

import (
	"github.com/labstack/echo/v4"
  )

func main()  {
  e := echo.New()
  e.GET("/hello", hello)
  e.Static("/", "public")
  e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {

  type Link struct {
    Url string
    Logo string
  }
  link := Link{Url: "https://facebook.com/sun", Logo: "TTTT"}

return c.JSON(200, link)
	// return c.String(http.StatusOK, "Hello, World! by Bunyapont Jongpremkitpaisan  :)")
}