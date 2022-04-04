package rest

type XForwardedAuthUser struct {
	UserId  string `json:"UserId"`
	GroupId string `json:"GroupId"`
	Viewer  string `json:"Viewer"`
}
