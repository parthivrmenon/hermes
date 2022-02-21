package link

type Link struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Hit struct {
	LINK_ID string `json:"id"`
	HITS    int    `json:"hits"`
}
