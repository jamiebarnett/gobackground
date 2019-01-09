package response

// type Images struct {
// 	Data []struct {
// 		ID            string        `json:"id"`
// 		Title         string        `json:"title"`
// 		Description   interface{}   `json:"description"`
// 		Datetime      int           `json:"datetime"`
// 		Type          string        `json:"type"`
// 		Animated      bool          `json:"animated"`
// 		Width         int           `json:"width"`
// 		Height        int           `json:"height"`
// 		Size          int           `json:"size"`
// 		Views         int           `json:"views"`
// 		Bandwidth     int64         `json:"bandwidth"`
// 		Vote          interface{}   `json:"vote"`
// 		Favorite      bool          `json:"favorite"`
// 		Nsfw          bool          `json:"nsfw"`
// 		Section       string        `json:"section"`
// 		AccountURL    interface{}   `json:"account_url"`
// 		AccountID     interface{}   `json:"account_id"`
// 		IsAd          bool          `json:"is_ad"`
// 		InMostViral   bool          `json:"in_most_viral"`
// 		HasSound      bool          `json:"has_sound"`
// 		Tags          []interface{} `json:"tags"`
// 		AdType        int           `json:"ad_type"`
// 		AdURL         string        `json:"ad_url"`
// 		InGallery     bool          `json:"in_gallery"`
// 		Link          string        `json:"link"`
// 		CommentCount  interface{}   `json:"comment_count"`
// 		FavoriteCount interface{}   `json:"favorite_count"`
// 		Ups           interface{}   `json:"ups"`
// 		Downs         interface{}   `json:"downs"`
// 		Points        interface{}   `json:"points"`
// 		Score         int           `json:"score"`
// 		IsAlbum       bool          `json:"is_album"`
// 	} `json:"data"`
// }

// Images stores all data returned from imgur API about an image
type Images struct {
	Data []struct {
		ID              string        `json:"id"`
		Title           string        `json:"title"`
		Description     interface{}   `json:"description"`
		Datetime        int           `json:"datetime"`
		Cover           string        `json:"cover"`
		CoverWidth      int           `json:"cover_width"`
		CoverHeight     int           `json:"cover_height"`
		AccountURL      interface{}   `json:"account_url"`
		AccountID       interface{}   `json:"account_id"`
		Privacy         interface{}   `json:"privacy"`
		Layout          interface{}   `json:"layout"`
		Views           int           `json:"views"`
		Link            string        `json:"link"`
		Ups             interface{}   `json:"ups"`
		Downs           interface{}   `json:"downs"`
		Points          interface{}   `json:"points"`
		Score           int           `json:"score"`
		IsAlbum         bool          `json:"is_album"`
		Vote            interface{}   `json:"vote"`
		Favorite        interface{}   `json:"favorite"`
		Nsfw            bool          `json:"nsfw"`
		Section         string        `json:"section"`
		CommentCount    interface{}   `json:"comment_count"`
		FavoriteCount   interface{}   `json:"favorite_count"`
		Topic           interface{}   `json:"topic"`
		TopicID         interface{}   `json:"topic_id"`
		ImagesCount     int           `json:"images_count"`
		InGallery       bool          `json:"in_gallery"`
		IsAd            bool          `json:"is_ad"`
		Tags            []interface{} `json:"tags"`
		AdType          int           `json:"ad_type"`
		AdURL           string        `json:"ad_url"`
		InMostViral     bool          `json:"in_most_viral"`
		IncludeAlbumAds bool          `json:"include_album_ads"`
		Images          []struct {
			ID            string        `json:"id"`
			Title         interface{}   `json:"title"`
			Description   interface{}   `json:"description"`
			Datetime      int           `json:"datetime"`
			Type          string        `json:"type"`
			Animated      bool          `json:"animated"`
			Width         int           `json:"width"`
			Height        int           `json:"height"`
			Size          int           `json:"size"`
			Views         int           `json:"views"`
			Bandwidth     int           `json:"bandwidth"`
			Vote          interface{}   `json:"vote"`
			Favorite      bool          `json:"favorite"`
			Nsfw          interface{}   `json:"nsfw"`
			Section       interface{}   `json:"section"`
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
			Score         interface{}   `json:"score"`
		} `json:"images"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
