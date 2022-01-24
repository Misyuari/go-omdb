package model

type Movie struct {
	Title      string          `json:"title,omitempty"`
	Year       string          `json:"year,omitempty"`
	Rated      string          `json:"rated,omitempty"`
	Released   string          `json:"released,omitempty"`
	Runtime    string          `json:"runtime,omitempty"`
	Genre      string          `json:"genre,omitempty"`
	Director   string          `json:"director,omitempty"`
	Writer     string          `json:"writer,omitempty"`
	Actors     string          `json:"actors,omitempty"`
	Plot       string          `json:"plot,omitempty"`
	Language   string          `json:"language,omitempty"`
	Country    string          `json:"country,omitempty"`
	Awards     string          `json:"awards,omitempty"`
	Poster     string          `json:"poster,omitempty"`
	Ratings    *[]MovieRatings `json:"ratings,omitempty"`
	Metascore  string          `json:"metascore,omitempty"`
	ImdbRating string          `json:"imdbRating,omitempty"`
	ImdbVotes  string          `json:"imdbVotes,omitempty"`
	ImdbID     string          `json:"imdbId,omitempty"`
	Type       string          `json:"type,omitempty"`
	DVD        string          `json:"dvd,omitempty"`
	BoxOffice  string          `json:"boxOffice,omitempty"`
	Production string          `json:"production,omitempty"`
	Website    string          `json:"website,omitempty"`
	Response   string          `json:"response,omitempty"`
}

type MovieRatings struct {
	Source string `json:"source,omitempty"`
	Value  string `json:"value,omitempty"`
}

type MovieRequestList struct {
	Page    string `json:"page" validate:"required,number,min=1"`
	Keyword string `json:"keyword" validate:"required"`
}
