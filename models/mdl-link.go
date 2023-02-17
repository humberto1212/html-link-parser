package models

type LinksList struct {
	List []LinksInfo
}

type LinksInfo struct {
	Href string
	Text string
}
