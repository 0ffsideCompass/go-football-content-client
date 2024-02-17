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
	Title    string               `json:"title"`
	Body     string               `json:"body"`
	Summary  string		      `json:"summary"`
	HashTags []string             `json:"hash_tags"`
	Type     string               `json:"type"`
	Data     FixtureDataStructure `json:"data"`
}

type FixtureByDateResponse struct {
	Data []LeagueFixtures
}

type LeagueFixtures struct {
	LeagueID   int    `json:"league_id"`
	LeagueName string `json:"league_name"`
	IDs        []int  `json:"ids"`
}

type FixtureDataStructure struct {
	League                 string `json:"league"`
	TeamOne                string `json:"team_one"`
	TeamTwo                string `json:"team_two"`
	TeamOneWins            int    `json:"team_one_wins"`
	TeamTwoWins            int    `json:"team_two_wins"`
	TeamOneDraws           int    `json:"team_one_draws"`
	TeamTwoDraws           int    `json:"team_two_draws"`
	TeamOneLoses           int    `json:"team_one_loses"`
	TeamTwoLoses           int    `json:"team_two_loses"`
	TeamOneGoals           int    `json:"team_one_goals"`
	TeamTwoGoals           int    `json:"team_two_goals"`
	TeamOneGoalsAgainst    int    `json:"team_one_goals_against"`
	TeamTwoGoalsAgainst    int    `json:"team_two_goals_against"`
	TeamOneCleanSheets     int    `json:"team_one_clean_sheets"`
	TeamTwoCleanSheets     int    `json:"team_two_clean_sheets"`
	TeamOneFailedToScore   int    `json:"team_one_failed_to_score"`
	TeamTwoFailedToScore   int    `json:"team_two_failed_to_score"`
	TeamOnePenaltiesScored int    `json:"team_one_penalties_scored"`
	TeamTwoPenaltiesScored int    `json:"team_two_penalties_scored"`
}
