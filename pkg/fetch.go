package osscontribute

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

const jsonFile = "data.json"

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	RepoURL     string `json:"repoURL"`
	AvatarURL   string `json:"avatarURL"`
	Language    string `json:"language"`
	Stars       int    `json:"stars"`
}

type Fetcher struct {
	Client *github.Client
	Cache  *Cache
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		Client: github.NewClient(nil),
		Cache:  NewCache(),
	}
}

func (f *Fetcher) RepoData(ctx context.Context) error {
	// check if there is existing data on disk to load from on boot.
	// only log if there is an error since this could be the first boot.
	fi, err := os.Stat(jsonFile)
	if err != nil {
		log.Println(err)
	}

	// cache JSON data if less than 24 hours old.
	// only log the error since this is a best effort attempt.
	if err == nil && time.Since(fi.ModTime()) < 24*time.Hour {
		data, err := os.ReadFile(jsonFile)
		if err != nil {
			log.Println(err)
		} else {
			f.Cache.Set(data)
			log.Println("Loaded existing, valid data into the cache. Skipping fetch...")
			return nil
		}
	}

	licenses, err := f.licenseKeys(ctx)
	if err != nil {
		return err
	}

	const baseQuery = "is:public archived:false good-first-issues:>=10 help-wanted-issues:>=10 stars:>=500"
	query := fmt.Sprintf("%s %s %s", baseQuery, licenses, oneMonthAgo())
	opts := &github.SearchOptions{}

	// Map to deduplicate repos
	// https://github.com/orgs/community/discussions/24361
	repoMap := make(map[string]Repo)

	log.Println("fetching GitHub repo data...")
	for page := 1; ; page++ {
		opts.Page = page
		opts.PerPage = 100

		result, resp, err := f.Client.Search.Repositories(ctx, query, opts)
		if err != nil {
			return fmt.Errorf("error searching repos: %w", err)
		}

		if len(result.Repositories) == 0 {
			break
		}

		for _, githubRepo := range result.Repositories {
			repo := Repo{
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

	var uniqueRepos []Repo
	for _, repo := range repoMap {
		uniqueRepos = append(uniqueRepos, repo)
	}

	// Sort repos by stars in descending order
	slices.SortStableFunc(uniqueRepos, func(a, b Repo) int {
		return b.Stars - a.Stars
	})

	data, err := json.MarshalIndent(uniqueRepos, "", "  ")
	if err != nil {
		return err
	}

	tmpfile, err := os.CreateTemp(".", "data-*.json")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(data); err != nil {
		return err
	}
	tmpfile.Close()

	if err := os.Rename(tmpfile.Name(), jsonFile); err != nil {
		return err
	}

	f.Cache.Set(data)

	log.Println("Successfully fetched and cached GitHub repo data")

	return nil
}

func (f *Fetcher) licenseKeys(ctx context.Context) (string, error) {
	licenses, _, err := f.Client.Licenses.List(ctx)
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
