package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type GithubHandler struct {
	flags        Flags
	githubClient *github.Client
	ctx          context.Context
}

func NewGithubHandler(flags Flags) (*GithubHandler, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: flags.GithubToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)

	githubClient, err := github.NewEnterpriseClient(NewGithubURL().API(flags.BaseURL), "", tc)
	if err != nil {
		return nil, err
	}
	return &GithubHandler{
		flags:        flags,
		githubClient: githubClient,
		ctx:          ctx,
	}, nil
}

func (h *GithubHandler) FetchCommits(owner, repo string, since, until time.Time) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{
		Author: h.flags.User,
		Since:  since,
		Until:  until,
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	var allCommits []*github.RepositoryCommit
	for {
		commits, resp, err := h.githubClient.Repositories.ListCommits(context.Background(), owner, repo, opt)
		if err != nil {
			return nil, err
		}
		allCommits = append(allCommits, commits...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allCommits, nil
}

// func (h *GithubHandler) FetchRepos() ([]*github.Repository, error) {
// 	var allRepos []*github.Repository
// 	opt := &github.RepositoryListOptions{
// 		ListOptions: github.ListOptions{
// 			PerPage: 100,
// 		},
// 	}

// 	for {
// 		repos, resp, err := h.githubClient.Repositories.List(context.Background(), h.flags.User, opt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		allRepos = append(allRepos, repos...)
// 		if resp.NextPage == 0 {
// 			break
// 		}
// 		opt.Page = resp.NextPage
// 	}

//		return allRepos, nil
//	}
func (h *GithubHandler) FetchRepos() ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 100}}

	// 1️⃣ Fetch user repositories
	for {
		repos, resp, err := h.githubClient.Repositories.List(context.Background(), "", opt) // "" -> Authenticated User
		if err != nil {
			return nil, fmt.Errorf("error fetching user repos: %w", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	// 2️⃣ Fetch user's organizations
	orgs, _, err := h.githubClient.Organizations.List(context.Background(), "", nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching user organizations: %w", err)
	}

	// 3️⃣ Fetch repositories for each organization
	for _, org := range orgs {
		orgOpt := &github.RepositoryListByOrgOptions{
			Type:        "all",
			ListOptions: github.ListOptions{PerPage: 100},
		}

		for {
			repos, resp, err := h.githubClient.Repositories.ListByOrg(context.Background(), org.GetLogin(), orgOpt)
			if err != nil {
				return nil, fmt.Errorf("error fetching repos for org %s: %w", org.GetLogin(), err)
			}

			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			orgOpt.Page = resp.NextPage
		}
	}

	return allRepos, nil
}
