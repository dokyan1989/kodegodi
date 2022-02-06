package jsonplaceholder

import (
	"fmt"
	"strings"

	"github.com/dokyan1989/kodegodi/lib/httputil"
)

type Client interface {
	GetPostById(id uint64) (Post, error)
	ListAllPosts() ([]Post, error)
	CreatePost(params CreatePostParam) (Post, error)
	UpdatePost(id uint64, params UpdatePostParam) (Post, error)
	PatchPost(id uint64, params PatchPostParam) (Post, error)
	DeletePost(id uint64) error
	SearchPosts(params SearchPostsParam) ([]Post, error)
}

type client struct {
	url string
}

func New() *client {
	return &client{url: "https://jsonplaceholder.typicode.com"}
}

var _ Client = &client{}

const (
	pathPost = "posts"
)

func (c *client) GetPostById(id uint64) (Post, error) {
	getPostByIdUrl := fmt.Sprintf("%s/%s/%d", c.url, pathPost, id)
	res, err := httputil.Get(getPostByIdUrl)
	if err != nil {
		return Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	post := Post{}
	err = httputil.ReadBodyJSON(res, &post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c *client) ListAllPosts() ([]Post, error) {
	listAllPostsUrl := fmt.Sprintf("%s/%s", c.url, pathPost)
	res, err := httputil.Get(listAllPostsUrl)
	if err != nil {
		return []Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return []Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	posts := []Post{}
	err = httputil.ReadBodyJSON(res, &posts)
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}

type CreatePostParam struct {
	UserID uint64 `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (c *client) CreatePost(params CreatePostParam) (Post, error) {
	createPostUrl := fmt.Sprintf("%s/%s", c.url, pathPost)
	res, err := httputil.Post(createPostUrl, httputil.MIMETypeApplicationJSON, params)
	if err != nil {
		return Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	post := Post{}
	err = httputil.ReadBodyJSON(res, &post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

type UpdatePostParam struct {
	UserID uint64 `json:"userId"`
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (c *client) UpdatePost(id uint64, params UpdatePostParam) (Post, error) {
	updatePostUrl := fmt.Sprintf("%s/%s/%d", c.url, pathPost, id)
	res, err := httputil.Put(updatePostUrl, httputil.MIMETypeApplicationJSON, params)
	if err != nil {
		return Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	post := Post{}
	err = httputil.ReadBodyJSON(res, &post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

type PatchPostParam struct {
	UserID uint64 `json:"userId,omitempty"`
	ID     uint64 `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func (c *client) PatchPost(id uint64, params PatchPostParam) (Post, error) {
	patchPostUrl := fmt.Sprintf("%s/%s/%d", c.url, pathPost, id)
	res, err := httputil.Patch(patchPostUrl, httputil.MIMETypeApplicationJSON, params)
	if err != nil {
		return Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	post := Post{}
	err = httputil.ReadBodyJSON(res, &post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c *client) DeletePost(id uint64) error {
	deletePostUrl := fmt.Sprintf("%s/%s/%d", c.url, pathPost, id)
	res, err := httputil.Delete(deletePostUrl)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return httputil.ErrResponseFailed(res.StatusCode)
	}

	return nil
}

type SearchPostsParam struct {
	UserID uint64 `json:"userId,omitempty"`
}

func (c *client) SearchPosts(params SearchPostsParam) ([]Post, error) {
	searchPostsUrl := c.buildSearchPostsUrl(params)
	res, err := httputil.Get(searchPostsUrl)
	if err != nil {
		return []Post{}, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return []Post{}, httputil.ErrResponseFailed(res.StatusCode)
	}

	posts := []Post{}
	err = httputil.ReadBodyJSON(res, &posts)
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}

func (c *client) buildSearchPostsUrl(params SearchPostsParam) string {
	url := fmt.Sprintf("%s/%s", c.url, pathPost)
	urlValues := []string{}
	if params.UserID > 0 {
		urlValues = append(urlValues, fmt.Sprintf("userId=%d", params.UserID))
	}

	if len(urlValues) > 0 {
		url = fmt.Sprintf("%s?%s", url, strings.Join(urlValues, "&"))
	}

	return url
}
