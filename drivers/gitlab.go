package drivers

import (
	"encoding/json"
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
	"login":      "Username",
	"avatar_url": "Avatar",
	"name":       "FullName",
}

// GitlabAPIMap is the map for API endpoints
var GitlabAPIMap = map[string]string{
	"endpoint":      "https://gitlab.com/api/v4",
	"userEndpoint":  "/user",
	"emailEndpoint": "/user/emails",
}

// GitlabUserFn is a callback to parse additional fields for User
var GitlabUserFn = func(client *http.Client, u *structs.User) {
	// Used to parse the email from response
	type additionalEmail struct {
		Email string `json:"email"`
	}
	var email []additionalEmail

	// Email can be nil because of the "keep my email private" setting
	if u.Email == "<nil>" {
		// Retrieve email
		req, err := client.Get(GitlabAPIMap["endpoint"] + GitlabAPIMap["emailEndpoint"])
		if err != nil {
			return
		}

		defer req.Body.Close()
		err = json.NewDecoder(req.Body).Decode(&email)
		if err != nil {
			return
		}

		u.Email = email[0].Email
	}
}

// GitlabDefaultScopes contains the default scopes
var GitlabDefaultScopes = []string{"user:email"}
