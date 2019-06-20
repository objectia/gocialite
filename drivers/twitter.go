package drivers

import (
	"fmt"
	"net/http"

	"github.com/objectia/gocialite/structs"
	"golang.org/x/oauth2"
)

const twitterDriverName = "twitter"

func init() {
	registerDriver(twitterDriverName, TwitterDefaultScopes, TwitterUserFn, TwitterEndpoint, TwitterAPIMap, TwitterUserMap)
}

// TwitterEndpoint is the oAuth endpoint
var TwitterEndpoint = oauth2.Endpoint{
	AuthURL:  "https://api.twitter.com/oauth/authorize",
	TokenURL: "https://api.twitter.com/oauth/access_token",
}

// TwitterUserMap is the map to create the User struct
var TwitterUserMap = map[string]string{}

// TwitterAPIMap is the map for API endpoints
var TwitterAPIMap = map[string]string{
	"endpoint":     "https://api.twitter.com",
	"userEndpoint": "/1.1/account/verify_credentials.json?include_email=true",
}

// TwitterUserFn is a callback to parse additional fields for User
var TwitterUserFn = func(client *http.Client, u *structs.User) {
	userData := u.Raw["data"].(map[string]interface{})
	u.ID = fmt.Sprintf("%.0f", userData["id"].(float64))
	u.FullName = userData["name"].(string)
	u.Avatar = userData["profile_image_url_https"].(string)

	u.Email = userData["email"].(string) // May be nil
}

// TwitterDefaultScopes contains the default scopes
var TwitterDefaultScopes = []string{}
