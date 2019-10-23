package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

var clientID = "6e33c9212621230df631"
var clientSecret = "e280813c8d831c9b9e0218f9bd750558a25a4bed"

// GithubHandler handles /github routes
type GithubHandler struct {
	http.ServeMux
}

// AccessTokenResponse represents the JSON object returned from GitHub's 'access token' request URL
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
}

func errorResponse(w http.ResponseWriter, msg string, err error) {
	log.Println(msg)
	log.Println(err)
	w.Write([]byte(msg))
	return
}

func convertCodeToAccessToken(w http.ResponseWriter, r *http.Request) {
	log.Println("[GithubHandler] hit /auth/rxcode")
	responseBody := []byte{}

	// Get query params
	code, _ := r.URL.Query()["code"]
	state, _ := r.URL.Query()["state"]
	log.Printf("[rxcode] code=%v, state=%v\n", code, state)
	
	// Prepare 'access token' URL
	accessTokenURL, _ := url.Parse("https://github.com/login/oauth/access_token")
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code[0])
	params.Add("client_id", clientID)
	params.Add("client_secret", clientSecret)
	accessTokenURL.RawQuery = params.Encode()

	// Prepare 'access token' request
	request, err := http.NewRequest("GET", accessTokenURL.String(), nil)
	request.Header.Add("Accept", "application/json")

	// Invoke 'access token' request
	log.Printf("[rxcode] making request for token -> %v\n", accessTokenURL.String())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		errorResponse(w, "ERROR getting access token...", err)
		return
	}

	log.Printf("[rxcode] response: %v - %v", response.StatusCode, response.Status)

	// Process response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorResponse(w, "ERROR unpacking response body...", err)
		return
	}
	msg := fmt.Sprintf("[rxcode] got token response -> %v\n", string(body))
	log.Printf(msg)
	responseBody = append(responseBody, msg...)

	// Interpret response as JSON data
	var accessTokenResponse AccessTokenResponse
	json.Unmarshal(body, &accessTokenResponse)
	msg = fmt.Sprintf("[rxcode] json response -> %v\n", accessTokenResponse)
	log.Printf(msg)
	responseBody = append(responseBody, msg...)
	msg = fmt.Sprintf("[rxcode] access token is: %v\n", accessTokenResponse.AccessToken)
	log.Printf(msg)
	responseBody = append(responseBody, msg...)

	// Provide HTTP Response
	http.SetCookie(w, &http.Cookie{Name: "LoginMarker", Value: accessTokenResponse.AccessToken, Domain: "rconway.co.uk", Path: "/"})
	w.Write(responseBody)

	// http.Redirect(w, r, "https://app.rconway.co.uk/oauth.html", 307)
}

// NewGithubHandler instantiates and returns a GithubHandler instance
func NewGithubHandler() http.Handler {
	handler := &GithubHandler{}

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[GithubHandler] hit /")
		w.Write([]byte("[GithubHandler] hit /"))
	})

	handler.HandleFunc("/auth/rxcode", convertCodeToAccessToken)

	return handler
}
