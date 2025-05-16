package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"gopkg.in/go-playground/pool.v3"
	"net/http"
)

type (
	User struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	Todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

func main() {
	p := pool.New()
	defer p.Close()
	userJob := p.Queue(fetchUser(1))
	todoJob := p.Queue(fetchTodo(1))
	userJob.Wait()
	if err := userJob.Error(); err != nil {
		panic(err)
	}
	todoJob.Wait()
	if err := todoJob.Error(); err != nil {
		panic(err)
	}
	user := userJob.Value().(*User)
	todo := todoJob.Value().([]*Todo)
	spew.Dump(user, todo)
}

func fetchUser(userId int) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		user := new(User)
		resp, err := resty.New().R().
			SetResult(user).
			Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userId))
		if err != nil {
			return nil, err
		}
		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("http %d error: %s", resp.StatusCode(), resp.Status())
		}
		return user, nil
	}
}

func fetchTodo(userId int) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		todos := make([]*Todo, 0)
		resp, err := resty.New().R().
			SetResult(&todos).
			Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d/todos", userId))
		if err != nil {
			return nil, err
		}
		if resp.StatusCode() != http.StatusOK {
			return nil, fmt.Errorf("http %d error: %s", resp.StatusCode(), resp.Status())
		}
		return todos, nil
	}
}
