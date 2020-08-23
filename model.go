package main

type URIDictionary struct {
	title string
}

type Bookmark struct {
	URLString     string
	URIDictionary map[string]string
}

type Folder struct {
	Title    string
	Children []Bookmark
}

type Root struct {
	Children []Folder
	Title    string
}
