package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kevincobain2000/wdhd/pkg"
)

var (
	flags   pkg.Flags
	version = "dev"
)

func main() {
	pkg.ParseFlags(&flags)
	if flags.Version {
		fmt.Println(version)
		return
	}
	if err := pkg.ValidateFlags(&flags); err != nil {
		log.Fatalf("Error validating flags: %v", err)
	}

	githubHandler, err := pkg.NewGithubHandler(flags)
	if err != nil {
		log.Fatalf("Error creating GitHub handler: %v", err)
	}

	repos, err := githubHandler.FetchRepos()
	if err != nil {
		log.Fatalf("Error fetching repos: %v", err)
	}

	fmt.Println(flags.Prompt)

	counter := 0

	fmt.Printf("From:%s, To:%s\n", flags.FromDate.Format("2006-01-02"), flags.ToDate.Format("2006-01-02"))
	for _, repo := range repos {
		owner := repo.GetOwner().GetLogin()
		repoName := repo.GetName()
		commits, err := githubHandler.FetchCommits(owner, repoName, flags.FromDate, flags.ToDate)
		if err != nil {
			continue
		}
		if len(commits) == 0 {
			continue
		}
		counter++
		fmt.Printf("\n%d) Contributions on org:%s, repo:%s\n", counter, owner, repoName)

		var prevMessage string
		for _, commit := range commits {
			message := commit.GetCommit().GetMessage()
			message = removeEmptyLines(message)
			if message == prevMessage {
				continue
			}
			fmt.Printf("- Commits on %s: %s\n", commit.GetCommit().GetCommitter().GetDate().Format("2006-01-02 15:04:05"), message)

			prevMessage = message
		}

		comments, prComments, err := githubHandler.FetchUserComments(owner, repoName)
		if err != nil {
			log.Printf("Error fetching comments for org:%s, repo:%s: %v", owner, repoName, err)
			continue
		}

		if len(comments) > 0 || len(prComments) > 0 {
			fmt.Printf("\nComments on org:%s, repo:%s\n", owner, repoName)
			for _, comment := range comments {
				fmt.Printf("- Issue Comment: %s\n", comment.GetBody())
			}
			for _, prComment := range prComments {
				fmt.Printf("- PR Comment: %s\n", prComment.GetBody())
			}
		}
	}
}

func removeEmptyLines(s string) string {
	lines := strings.Split(s, "\n")
	var filtered []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filtered = append(filtered, line)
		}
	}
	return strings.Join(filtered, "\n")
}
