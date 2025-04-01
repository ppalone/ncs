package ncs

// Artist
type Artist struct {
	Id        string
	Name      string
	ArtistURL string
}

// ArtistInfo
type ArtistInfo struct {
	Id         string
	Name       string
	CoverImage string
	Genres     []string
	Songs      []Song
}
