package pr

import "os"

var (
	token = os.Getenv("GITHUB_TOKEN")
)

func SetToken(t string) {
	token = t
}
