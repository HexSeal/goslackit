package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	api := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}

    user, err := api.GetUserInfo("U023BECGF")
    if err != nil {
		fmt.Printf("%s\n", err)
		return
    }
    fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}