package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		assert.FailNow(t, "Failed to make request")
	}
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	result := make(map[string]interface{})
	json.Unmarshal(body, &result)
	fmt.Println(result["id"], result["title"], result["completed"])
	assert.Equal(t, float64(1), result["id"])
	assert.Equal(t, "delectus aut autem", result["title"])
	assert.Equal(t, false, result["completed"])
}

func TestPostRequest(t *testing.T) {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", nil)
	if err != nil {
		assert.FailNow(t, "Failed to make request")
	}
	defer resp.Body.Close()
	assert.Equal(t, 201, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	result := make(map[string]interface{})
	json.Unmarshal(body, &result)
	assert.Equal(t, float64(101), result["id"])
}

func TestPutRequest(t *testing.T) {
	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		assert.FailNow(t, "Failed to create request")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		assert.FailNow(t, "Failed to make request")
	}
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	result := make(map[string]interface{})
	json.Unmarshal(body, &result)
	assert.Equal(t, float64(1), result["id"])
}

func TestDeleteRequest(t *testing.T) {
	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		assert.FailNow(t, "Failed to create request")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		assert.FailNow(t, "Failed to make request")
	}
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}
