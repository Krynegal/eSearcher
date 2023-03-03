package handlers

import (
	"eSearcher/configs"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type GoogleAuth struct {
	OauthConfGl        *oauth2.Config
	oauthStateStringGl string
}

func NewGoogleAuth(cfg *configs.Config) *GoogleAuth {
	oauthConfGl := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://%s:%s/%s", cfg.ServerHost, cfg.ServerPort, cfg.RedirectPath),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &GoogleAuth{
		OauthConfGl:        oauthConfGl,
		oauthStateStringGl: "",
	}
}

func (r *Router) HandleMain(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}

func HandleLogin(w http.ResponseWriter, req *http.Request, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Printf("Parse: " + err.Error())
	}
	log.Println(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	log.Println(url)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (r *Router) HandleGoogleLogin(w http.ResponseWriter, req *http.Request) {
	HandleLogin(w, req, r.GoogleAuth.OauthConfGl, r.GoogleAuth.oauthStateStringGl)
}

func (r *Router) CallBackFromGoogle(w http.ResponseWriter, req *http.Request) {
	log.Println("Callback-gl..")

	state := req.FormValue("state")
	log.Println(state)
	if state != r.GoogleAuth.oauthStateStringGl {
		log.Println("invalid oauth state, expected " + r.GoogleAuth.oauthStateStringGl + ", got " + state + "\n")
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
		return
	}

	code := req.FormValue("code")
	log.Println(code)

	if code == "" {
		log.Println("Code not found..")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := req.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := r.GoogleAuth.OauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Printf("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		log.Println("TOKEN>> AccessToken>> " + token.AccessToken)
		log.Println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		log.Println("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			log.Printf("Get: " + err.Error() + "\n")
			http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("ReadAll: " + err.Error() + "\n")
			http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
			return
		}

		log.Printf("parseResponseBody: " + string(response) + "\n")

		w.Write([]byte("Hello, I'm protected\n"))
		w.Write(response)
		return
	}
}
