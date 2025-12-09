package aoc

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

var sessionStorage = "session"
var baseUrl = "https://adventofcode.com"

func acquireSessionCookie() *http.Cookie {
	file, err := os.ReadFile(sessionStorage)
	cookie := &http.Cookie{
		Name: "session",
	}

	if err == nil {
		cookie.Value = string(file)
		return cookie
	}

	session := login()
	_ = os.WriteFile(sessionStorage, []byte(session), 0644)
	cookie.Value = session
	return cookie
}

func createHttpClient(sessionCookie *http.Cookie) *http.Client {
	jar, _ := cookiejar.New(nil)
	var parsedUrl, _ = url.Parse(baseUrl)
	jar.SetCookies(parsedUrl, []*http.Cookie{sessionCookie})
	client := &http.Client{
		Jar: jar,
	}
	return client
}

type Client struct {
	LoadInput func(year int, day int) string
}

func CreateClient() Client {
	httpClient := createHttpClient(acquireSessionCookie())
	var loadInput = func(year int, day int) string {
		res, err := httpClient.Get(fmt.Sprintf("%s/%d/day/%d/input", baseUrl, year, day))
		if err != nil {
			panic(err)
		}
		body, _ := io.ReadAll(res.Body)
		return string(body)
	}

	return Client{
		LoadInput: loadInput,
	}
}
