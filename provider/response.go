package provider

// Response struct to parse the API response query of
// http://ow-api.herokuapp.com/profile/pc/us/[BATTLETAG]
type Response struct {
	Username string `json:"username"`
	Games    struct {
		Competitive struct {
			Won    int `json:"won"`
			Lost   int `json:"lost"`
			Draw   int `json:"draw"`
			Played int `json:"played"`
		} `json:"competitive"`
	} `json:"games"`
	Competitive struct {
		Rank    int    `json:"rank"`
		RankImg string `json:"rank_img"`
	} `json:"competitive"`
}
