package drivers

import (
	"net/http"

	"github.com/objectia/gocialite/structs"
	"golang.org/x/oauth2/gitlab"
)

const gitlabDriverName = "gitlab"

func init() {
	registerDriver(gitlabDriverName, GitlabDefaultScopes, GitlabUserFn, gitlab.Endpoint, GitlabAPIMap, GitlabUserMap)
}

// GitlabUserMap is the map to create the User struct
var GitlabUserMap = map[string]string{
	"id":         "ID",
	"email":      "Email",
	"username":   "Username",
	"avatar_url": "Avatar",
	"name":       "FullName",
}

// GitlabAPIMap is the map for API endpoints
var GitlabAPIMap = map[string]string{
	"endpoint":     "https://gitlab.com/api/v4",
	"userEndpoint": "/user",
}

// GitlabUserFn is a callback to parse additional fields for User
var GitlabUserFn = func(client *http.Client, u *structs.User) {}

// GitlabDefaultScopes contains the default scopes
var GitlabDefaultScopes = []string{"read_user"}
