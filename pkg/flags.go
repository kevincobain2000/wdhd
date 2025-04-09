package pkg

import (
	"flag"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	prompt = `
Summarize the following commit messages for evaluation purposes
The format is: #) Contributions on org:<org>, repo:<repo>
The format for each commit is <date>: <message>

Summarize contributions on each org and repo in about 1-2 sentences each to highlight the key contributions
How smart was the developer, how much did they contribute, and how did they help the team?
How much %age of his time was spent overall?
How many people did he help?
How many people did he collaborate with?
How many review comments did he make? And what were their impacts?
If there are merges, then count those as his tech lead contributions.
If there are bug fixes, then count those as his debugging contributions.
If there are new features, then count those as his feature contributions.
If there are refactorings, then count those as his code quality contributions.
Make judgements based on commit messages.
`
)

type Flags struct {
	// required
	GithubToken string // set by env var GITHUB_TOKEN
	User        string
	BaseURL     string
	Prompt      string
	DaysAgo     int
	FromDate    time.Time
	ToDate      time.Time
	Version     bool
}

const (
	githubTokenEnv = "GITHUB_TOKEN" // nolint: gosec
	defaultBaseURL = "github.com"
)

func ParseFlags(f *Flags) {
	flag.StringVar(&f.GithubToken, "token", "", "GITHUB_TOKEN via env or flag")
	flag.StringVar(&f.User, "user", "", "GitHub user")
	flag.StringVar(&f.BaseURL, "base-url", defaultBaseURL, "GitHub base URL")
	flag.StringVar(&f.Prompt, "prompt", prompt, "Prompt for the user")
	flag.IntVar(&f.DaysAgo, "days", 7, "Number of days ago to fetch commits")
	flag.BoolVar(&f.Version, "version", false, "print version and exit")

	flag.Parse()
}

func ValidateFlags(f *Flags) error {
	if f.GithubToken == "" {
		f.GithubToken = os.Getenv(githubTokenEnv)
	}
	var err error
	f.BaseURL, err = extractHost(f.BaseURL)
	if err != nil {
		return err
	}
	f.FromDate, f.ToDate = getDefaultDates(f.DaysAgo)
	if f.GithubToken == "" || f.User == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return nil
}

func getDefaultDates(daysAgo int) (time.Time, time.Time) {
	now := time.Now()
	from := now.AddDate(0, 0, -daysAgo).UTC()
	to := now.UTC()
	return from, to
}

func extractHost(fullURL string) (string, error) {
	if !strings.HasPrefix(fullURL, "http") {
		fullURL = "https://" + fullURL
	}
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	return parsedURL.Host, nil
}
