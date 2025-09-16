package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// создаём контекст с максимальным временем запроса
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(time.Millisecond*5000))
	defer cancel()
	// создаём запрос к серверу http://httpbin.org, ответ будет возвращаться
	// с задержкой в одну секунду, которая указана в запросе /delay/1
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://httpbin.org/delay/1", nil)
	// эквивалентно
	// req, err := http.NewRequest(http.MethodGet,
	//                          "http://httpbin.org/delay/1", nil)
	// ...
	// req.WithContext(ctx)

	if err != nil {
		log.Fatal(err)
	}
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
