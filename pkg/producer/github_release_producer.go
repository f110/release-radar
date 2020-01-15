package producer

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
	"golang.org/x/xerrors"
)

type GithubReleaseProducer struct {
	Org  string
	Repo string

	client *github.Client
	db     *DataSource
}

func NewGithubRelease(org, repo string, db *DataSource) *GithubReleaseProducer {
	return &GithubReleaseProducer{Org: org, Repo: repo, db: db, client: github.NewClient(nil)}
}

func (g *GithubReleaseProducer) Produce() ([]*Release, error) {
	releases, _, err := g.client.Repositories.ListReleases(context.Background(), g.Org, g.Repo, &github.ListOptions{PerPage: 100})
	if err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	fetchedReleases := make([]*Release, 0)
	for _, v := range releases {
		ver, err := ParseVersionString(v.GetTagName())
		if err != nil {
			continue
		}
		fetchedReleases = append(fetchedReleases, &Release{Version: ver, Published: v.GetPublishedAt().Time, Author: v.Author.GetLogin()})
	}

	oldReleases, err := g.db.List(fmt.Sprintf("github/%s/%s", g.Org, g.Repo))
	if err != nil {
		return nil, xerrors.Errorf(": %v", err)
	}

	return Releases(fetchedReleases).Diff(oldReleases), nil
}
