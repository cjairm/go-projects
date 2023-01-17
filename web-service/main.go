package main

import (
	albums "cjairm/web-service-gin/albums"

	"github.com/gin-gonic/gin"
)

func main() {
  // Initialize a Gin router using Default.
  router := gin.Default()
  // Use the GET function to associate the GET HTTP method
  // and /albums path with a handler function.
  router.GET("/albums", albums.GetAlbums)
  // Associate the POST method at the /albums path with the
  // postAlbums function.
  router.POST("/albums", albums.PostAlbums)
  // Associate the /albums/:id path with the getAlbumByID
  // function. In Gin, the colon preceding an item in the path
  // signifies that the item is a path parameter.
  router.GET("/albums/:id", albums.GetAlbumById)

  // Start server
  router.Run("localhost:8010")
}

// go get .
// go run .
// http://localhost:8010/albums
