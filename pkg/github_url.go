package pkg

import (
	"fmt"
	"strings"
)

type GithubURL struct {
}

func NewGithubURL() *GithubURL {
	return &GithubURL{}
}

func (e *GithubURL) API(baseURL string) string {
	s := fmt.Sprintf("https://%s/api/v3", baseURL)
	return s
}

func (e *GithubURL) Web(baseURL string) string {
	s := fmt.Sprintf("https://%s", baseURL)
	return s
}

func (e *GithubURL) CloneURL(baseURL, orgName, repoName string) string {
	s := fmt.Sprintf("%s/%s/%s.git", e.Web(baseURL), orgName, repoName)
	return s
}

func (e *GithubURL) RemoteURL(baseURL, orgName, repoName, token string) string {
	host := e.Web(baseURL)
	host = strings.TrimPrefix(host, "https://")
	s := fmt.Sprintf("https://%s@%s/%s/%s.git", token, host, orgName, repoName)
	return s
}
