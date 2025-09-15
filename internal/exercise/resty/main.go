package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type MyApiError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// Post — модель, описание основного объекта.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	client := resty.New()
	_, err := client.R().
		SetResult(&users).
		Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	// если выбрали resty, используйте SetResult(&users)
	// для получения результата сразу в виде массива
	// ...
}

func post() {
	client := resty.New()

	var responseErr MyApiError
	var post Post

	_, err := client.R().
		SetError(&responseErr).
		SetResult(&post).
		Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println(responseErr)
		panic(err)
	}

	fmt.Printf("%#v\n", post)
}

func get() {
	// создаём новый клиент
	client := resty.New()

	resp, err := client.R().
		SetAuthToken("Bearer <TOKEN>").
		Get("https://www.yfull.com/static/img/fl/24/ru.png")

	fmt.Println("Исследуем объект Response:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println("----")
}
