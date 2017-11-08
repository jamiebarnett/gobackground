package response

type Images struct {
	Data []struct {
		ID            string        `json:"id"`
		Title         string        `json:"title"`
		Description   interface{}   `json:"description"`
		Datetime      int           `json:"datetime"`
		Type          string        `json:"type"`
		Animated      bool          `json:"animated"`
		Width         int           `json:"width"`
		Height        int           `json:"height"`
		Size          int           `json:"size"`
		Views         int           `json:"views"`
		Bandwidth     int64         `json:"bandwidth"`
		Vote          interface{}   `json:"vote"`
		Favorite      bool          `json:"favorite"`
		Nsfw          bool          `json:"nsfw"`
		Section       string        `json:"section"`
		AccountURL    interface{}   `json:"account_url"`
		AccountID     interface{}   `json:"account_id"`
		IsAd          bool          `json:"is_ad"`
		InMostViral   bool          `json:"in_most_viral"`
		HasSound      bool          `json:"has_sound"`
		Tags          []interface{} `json:"tags"`
		AdType        int           `json:"ad_type"`
		AdURL         string        `json:"ad_url"`
		InGallery     bool          `json:"in_gallery"`
		Link          string        `json:"link"`
		CommentCount  interface{}   `json:"comment_count"`
		FavoriteCount interface{}   `json:"favorite_count"`
		Ups           interface{}   `json:"ups"`
		Downs         interface{}   `json:"downs"`
		Points        interface{}   `json:"points"`
		Score         int           `json:"score"`
		IsAlbum       bool          `json:"is_album"`
	} `json:"data"`
}
