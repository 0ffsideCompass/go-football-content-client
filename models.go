package client

type ArticleGenerateRequest struct {
	ID          int    `json:"id"`
	ContentType string `json:"content_type"`
	Language    string `json:"language"`
}

type FixtureByDateRequest struct {
	Date string `json:"date"`
}

type ArticleResponse struct {
	Title    string      `json:"title"`
	Body     string      `json:"body"`
	HashTags []string    `json:"hash_tags"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}

type FixtureByDateResponse struct {
	Data []LeagueFixtures
}

type LeagueFixtures struct {
	LeagueID   int    `json:"league_id"`
	LeagueName string `json:"league_name"`
	IDs        []int  `json:"ids"`
}
