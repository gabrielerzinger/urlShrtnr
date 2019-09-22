package models

// Entry is the main entity in the system
type Entry struct {
	ShortUrl string
	URL string
}


type EntryRequest struct {
	URL string
}
