package shortUrl

type ShortUrl struct {
	ID        int    `json:"_id"` //bson是用来创建后返回，omitempty是创建时自动创建
	ShortUrl  string `json:"shortUrl"`
	OriginUrl string `json:"originUrl"`
	HashCode  string `json:"hashCode"`
}

type Long2ShortRequest struct {
	OriginUrl string `json:"origin_url"`
}

type ResponseHeader struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Long2ShortResponse struct {
	ResponseHeader
	ShortUrl string `json:"short_url"`
}

type Short2LongRequest struct {
	ShortUrl string `json:"short_url"`
}

type Short2LongResponse struct {
	ResponseHeader
	OriginUrl string `json:"origin_url"`
}
