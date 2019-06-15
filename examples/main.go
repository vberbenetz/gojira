package examples

import (
	"context"
	"log"
	"../gojira"
)

// Basic auth example with Jira Cloud API Key
func main() {
	ctx := context.Background()

	tp := gojira.BasicAuth{
		Username: "YOUR_JIRA_USER_NAME",
		ApiKey: "API_KEY_REGISTERED_WITH_YOUR_USER_NAME",
	}

	client, err := gojira.NewClient(tp.Client(), "YOUR_SUB_DOMAIN")

	log.Println(err)

	roles, _, err := client.ApplicationRole.List(ctx)

	log.Println(roles)
	log.Println(err)
}
