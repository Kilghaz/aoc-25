package aoc

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func login() string {
	var authUrl = "https://adventofcode.com/auth/github"
	var port = 8001
	var chromeDriverPath = "/opt/homebrew/bin/chromedriver"

	var opts []selenium.ServiceOption
	service, _ := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)

	defer func(service *selenium.Service) {
		_ = service.Stop()
	}(service)

	capabilities := selenium.Capabilities{"browserName": "chrome"}
	webDriver, _ := selenium.NewRemote(capabilities, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	defer func(webDriver selenium.WebDriver) {
		_ = webDriver.Quit()
	}(webDriver)

	_ = webDriver.Get(authUrl)

	for {
		cookies, _ := webDriver.GetCookies()
		for _, cookie := range cookies {
			if cookie.Name == "session" {
				return cookie.Value
			}
		}

		time.Sleep(1 * time.Second)
	}
}
