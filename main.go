package main

import (
	"flag"
)

func main() {
	bookmarkFilePath := flag.String("bookmarksFilePath", DEFAULT_BOOKMARK_FILE_PATH, "Path to Bookmarks.plist file")
	serverPort := flag.Int("serverPort", DEFAULT_SERVER_PORT, "Server port to listen")
	flag.Parse()
	BookmarkFilePath = *bookmarkFilePath
	ServerPort = *serverPort
	Serve()
}
