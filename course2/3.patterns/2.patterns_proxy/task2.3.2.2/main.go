package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "your_access_token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	g := NewGithubAdapter(client)

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "ptflp"))
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}

	return g
}

type GithubProxy struct {
	githuber Githuber
	cache    map[string][]Item
}

func NewGithubProxy(githubClient *github.Client) *GithubProxy {
	gp := &GithubProxy{
		githuber: NewGithubAdapter(githubClient),
		cache:    map[string][]Item{},
	}

	return gp
}

func (gp GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	if value, ok := gp.cache[username]; ok {
		return value, nil
	}

	items, err := gp.githuber.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}

	gp.cache[username] = items

	return items, nil
}

func (gp GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	if value, ok := gp.cache[username]; ok {
		return value, nil
	}

	items, err := gp.githuber.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}

	gp.cache[username] = items

	return items, nil
}

type Item struct {
	Title       string
	Description string
	Link        string
}

func (ga GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := ga.GistList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0)
	for _, gist := range gists {
		item := Item{
			Title:       *gist.ID,
			Description: *gist.Description,
			Link:        *gist.HTMLURL,
		}
		items = append(items, item)
	}
	return items, nil
}

func (ga GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := ga.RepoList.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, repo := range repos {
		item := Item{
			Title:       *repo.FullName,
			Description: *repo.Description,
			Link:        *repo.HTMLURL,
		}
		items = append(items, item)
	}

	return items, nil
}
