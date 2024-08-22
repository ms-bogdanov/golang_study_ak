package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

type Item struct {
	Title       string
	Description string
	Link        string
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GithubGist struct {
	client *github.Client
}

func NewGithubGist(client *github.Client) *GithubGist {
	return &GithubGist{
		client: client,
	}
}

type GithubRepo struct {
	client *github.Client
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	return &GithubRepo{
		client: client,
	}
}

func (gg GithubRepo) GetItems(ctx context.Context, username string) (items []Item, err error) {
	repos, _, err := gg.client.Repositories.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

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

func (gg GithubGist) GetItems(ctx context.Context, username string) (items []Item, err error) {
	gists, _, err := gg.client.Gists.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}

	items = make([]Item, 0)
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

type GeneralGithub struct {
	client *github.Client
}

func NewGeneralGithub(client *github.Client) *GeneralGithub {
	return &GeneralGithub{
		client: client,
	}
}

func (gg GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) (items []Item, err error) {
	return strategy.GetItems(ctx, username)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "your-access-token"})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	data, err = gg.GetItems(context.Background(), "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
