package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	clientRead()
}

func postLocalhostFormUrlEncoded() {
	URL := "http://localhost:8080"
	// готовим контейнер для данных
	// используем тип url.Values из пакета net/url
	data := url.Values{}
	// устанавливаем данные
	data.Set("key1", "value1")
	data.Set("key2", "value2")
	// пишем запрос
	request, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(data.Encode()))
	if err != nil {
		// обрабатываем ошибку
	}
	// устанавливаем заголовки
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, response.Body) // вывод!!! ответа в консоль
	response.Body.Close()
}

func postLocalhostMultipartFormData() {
	filename := "file.txt"
	file, _ := os.Open(filename) // открываем файл
	defer file.Close()           // не забываем закрыть
	body := &bytes.Buffer{}      // создаём буфер
	// на основе буфера конструируем multipart.Writer из пакета mime/multipart
	writer := multipart.NewWriter(body)
	// готовим форму для отправки файла на сервер
	part, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		// обрабатываем ошибку
	}
	// копируем файл в форму
	// multipart.Writer отформатирует данные и запишет в предоставленный буфер
	_, err = io.Copy(part, file)
	if err != nil {
		// обрабатываем ошибку
	}
	writer.Close()

	url := "http://localhost:8080"
	client := &http.Client{}
	// пишем запрос
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		// обрабатываем ошибку
	}
	// добавляем заголовок запроса
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, response.Body) // вывод!!! ответа в консоль
	response.Body.Close()
}

func postLocalhostJson() {
	url := "http://localhost:8080"
	client := &http.Client{}
	var body = []byte(`{"message":"Hello"}`)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, response.Body) // вывод!!! ответа в консоль
	response.Body.Close()
}

func getLocalhost() {
	url := "http://localhost:8080"
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, response.Body) // вывод!!! ответа в консоль
	response.Body.Close()
}

func clientRead() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
}

func httpRead() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	if _, err = io.CopyN(os.Stdout, response.Body, 512); err != nil {
		fmt.Println(err)
	}
	/* другой вариант
	    body, err := io.ReadAll(response.Body)
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	   if len(body) > 512 {
	        body = body[:512]
	    }
	    fmt.Print(string(body)) */
}
