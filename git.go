/*
 * Copyright © 2024 weidongkl <weidongkx@gmail.com>
 */

package git

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"io"
	"os"
	"path"
	"strings"
)

const (
	suffix = ".git"
)

type clone struct {
	url      string
	branch   string
	username string
	password string
	depth    int
	progress io.Writer
	dstDir   string
}

func NewClone(url string) *clone {
	urlFields := strings.Split(url, "/")
	baseName := urlFields[len(urlFields)-1]
	if strings.HasSuffix(baseName, suffix) {
		baseName = strings.TrimSuffix(baseName, suffix)
	}
	return &clone{url: url,
		//depth:    1,
		progress: os.Stdout,
		dstDir: path.Join(os.TempDir(),
			baseName)}
}

func (c *clone) SetAuth(username, password string) *clone {
	c.username = username
	c.password = password
	return c
}

func (c *clone) SetToken(token string) *clone {
	return c.SetAuth("abc123", token)
}

func (c *clone) SetDepth(depth int) *clone {
	c.depth = depth
	return c
}

func (c *clone) SetBranch(branch string) *clone {
	c.branch = branch
	return c
}

func (c *clone) SetProgress(reader io.Writer) *clone {
	c.progress = reader
	return c
}

func (c *clone) SetDstDir(dir string) *clone {
	c.dstDir = dir
	return c
}

func (c *clone) Clone() error {
	cloneOptions := &git.CloneOptions{
		URL: c.url,
	}
	if c.branch != "" {
		cloneOptions.ReferenceName = plumbing.NewBranchReferenceName(c.branch)
	}
	if c.username != "" {
		cloneOptions.Auth = &http.BasicAuth{
			Username: c.username,
			Password: c.password,
		}
	}
	if c.depth > 0 {
		cloneOptions.Depth = c.depth
	}

	if c.progress != nil {
		cloneOptions.Progress = c.progress
	}

	_, err := git.PlainClone(c.dstDir, false, cloneOptions)
	return err
}

type pull struct {
	path     string
	username string
	password string
	progress io.Writer
}

func NewPull(path string) *pull {
	return &pull{path: path, progress: os.Stdout}
}

func (c *pull) SetAuth(username, password string) *pull {
	c.username = username
	c.password = password
	return c
}

func (c *pull) SetToken(token string) *pull {
	return c.SetAuth("abc123", token)
}

func (c *pull) SetProgress(progress io.Writer) *pull {
	c.progress = progress
	return c
}

func (c *pull) Pull() error {
	pullOptions := &git.PullOptions{
		RemoteName: "origin",
	}
	if c.username != "" {
		pullOptions.Auth = &http.BasicAuth{
			Username: c.username,
			Password: c.password,
		}
	}
	r, err := git.PlainOpen(c.path)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(pullOptions)
	if err == git.ErrNonFastForwardUpdate {
		return nil
	}
	return err
}
