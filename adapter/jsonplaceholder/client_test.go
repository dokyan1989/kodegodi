package jsonplaceholder

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dokyan1989/kodegodi/lib/httputil"
)

func Test_client_GetPostById_Success(t *testing.T) {
	want := Post{
		UserID: 1,
		ID:     1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusOK)
	}))

	c := &client{url: mockServer.URL}
	post, err := c.GetPostById(1)
	if err != nil {
		t.Errorf("client.GetPostById() error = %v", err)
		return
	}
	if !reflect.DeepEqual(post, want) {
		t.Errorf("client.GetPostById() = %v, want %v", post, want)
	}
}

func Test_client_ListAllPosts_Success(t *testing.T) {
	want := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		},
		{
			UserID: 1,
			ID:     2,
			Title:  "qui est esse",
			Body:   "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusOK)
	}))

	c := &client{url: mockServer.URL}
	posts, err := c.ListAllPosts()
	if err != nil {
		t.Errorf("client.ListAllPosts() error = %v", err)
		return
	}
	if !reflect.DeepEqual(posts, want) {
		t.Errorf("client.ListAllPosts() = %v, want %v", posts, want)
	}
}

func Test_client_CreatePost_Success(t *testing.T) {
	want := Post{
		UserID: 1,
		ID:     1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusCreated)
	}))
	c := &client{url: mockServer.URL}
	createdPost, err := c.CreatePost(CreatePostParam{
		UserID: 1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	})
	if err != nil {
		t.Errorf("client.CreatePost() error = %v", err)
		return
	}
	if !reflect.DeepEqual(createdPost, want) {
		t.Errorf("client.CreatePost() = %v, want %v", createdPost, want)
	}
}

func Test_client_UpdatePost_Success(t *testing.T) {
	want := Post{
		UserID: 1,
		ID:     1,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusOK)
	}))
	c := &client{url: mockServer.URL}
	createdPost, err := c.UpdatePost(
		1,
		UpdatePostParam{
			UserID: 1,
			ID:     1,
			Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		},
	)
	if err != nil {
		t.Errorf("client.UpdatePost() error = %v", err)
		return
	}
	if !reflect.DeepEqual(createdPost, want) {
		t.Errorf("client.UpdatePost() = %v, want %v", createdPost, want)
	}
}

func Test_client_PatchPost_Success(t *testing.T) {
	want := Post{
		UserID: 1,
		ID:     1,
		Title:  "foo",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusOK)
	}))
	c := &client{url: mockServer.URL}
	updatedPost, err := c.PatchPost(1, PatchPostParam{Title: "foo"})
	if err != nil {
		t.Errorf("client.PatchPost() error = %v", err)
		return
	}
	if !reflect.DeepEqual(updatedPost, want) {
		t.Errorf("client.PatchPost() = %v, want %v", updatedPost, want)
	}
}

func Test_client_DeletePost_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, 1, http.StatusOK)
	}))
	c := &client{url: mockServer.URL}
	err := c.DeletePost(1)
	if err != nil {
		t.Errorf("client.DeletePost() error = %v", err)
		return
	}
}

func Test_client_SearchPosts_Success(t *testing.T) {
	want := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		},
		{
			UserID: 1,
			ID:     2,
			Title:  "qui est esse",
			Body:   "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httputil.WriteJSON(w, want, http.StatusOK)
	}))

	c := &client{url: mockServer.URL}
	posts, err := c.SearchPosts(SearchPostsParam{UserID: 1})
	if err != nil {
		t.Errorf("client.SearchPosts() error = %v", err)
		return
	}
	if !reflect.DeepEqual(posts, want) {
		t.Errorf("client.SearchPosts() = %v, want %v", posts, want)
	}
}
