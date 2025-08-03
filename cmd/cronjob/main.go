package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/go-github/v64/github"
)

const (
	jsonFile = "data.json"
	gistID   = "dafa982abfa42982e02c75f1ddec46be"
)

type repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	RepoURL     string `json:"repoURL"`
	AvatarURL   string `json:"avatarURL"`
	Language    string `json:"language"`
	Stars       int    `json:"stars"`
}

func main() {
	ctx := context.Background()

	token := os.Getenv("GIST_TOKEN")
	if token == "" {
		log.Fatal("GIST_TOKEN environment variable is required")
	}

	client := github.NewClient(nil).WithAuthToken(token)

	licenses, err := licenseKeys(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	const baseQuery = "is:public archived:false good-first-issues:>=10 help-wanted-issues:>=10 stars:>=500"
	query := fmt.Sprintf("%s %s %s", baseQuery, licenses, oneMonthAgo())
	opts := &github.SearchOptions{}

	// Map to deduplicate repos
	// https://github.com/orgs/community/discussions/24361
	repoMap := make(map[string]repo)

	log.Println("fetching GitHub repo data...")
	for page := 1; ; page++ {
		opts.Page = page
		opts.PerPage = 100

		result, resp, err := client.Search.Repositories(ctx, query, opts)
		if err != nil {
			log.Fatalf("error searching repos: %v", err)
		}
		if len(result.Repositories) == 0 {
			log.Fatal("unexpected error: no GitHub repositories found matching the specified search criteria")
		}

		for _, githubRepo := range result.Repositories {
			repo := repo{
				Name:        githubRepo.GetName(),
				Description: githubRepo.GetDescription(),
				Owner:       *githubRepo.Owner.Login,
				RepoURL:     githubRepo.GetHTMLURL(),
				AvatarURL:   githubRepo.Owner.GetAvatarURL(),
				Language:    githubRepo.GetLanguage(),
				Stars:       githubRepo.GetStargazersCount(),
			}
			repoMap[repo.Name] = repo
		}

		if resp.NextPage == 0 {
			break
		}
	}

	var uniqueRepos []repo
	for _, repo := range repoMap {
		uniqueRepos = append(uniqueRepos, repo)
	}

	// Sort repos by stars in descending order
	slices.SortStableFunc(uniqueRepos, func(a, b repo) int {
		return b.Stars - a.Stars
	})

	data, err := json.MarshalIndent(uniqueRepos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	gist := &github.Gist{
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(jsonFile): {
				Content: github.String(string(data)),
			},
		},
	}

	_, _, err = client.Gists.Edit(ctx, gistID, gist)
	if err != nil {
		log.Fatalf("error updating gist: %v", err)
	}

	log.Printf("Successfully fetched and stored GitHub repo data to gist: %s", gistID)
}

func licenseKeys(ctx context.Context, client *github.Client) (string, error) {
	licenses, _, err := client.Licenses.List(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to list licenses: %w", err)
	}

	var licenseKeys []string
	for _, license := range licenses {
		licenseKeys = append(licenseKeys, fmt.Sprintf("license:%s", license.GetKey()))
	}
	return strings.Join(licenseKeys, " "), nil
}

func oneMonthAgo() string {
	return fmt.Sprintf("pushed:>%s", time.Now().AddDate(0, -1, 0).Format(time.DateOnly))
}
