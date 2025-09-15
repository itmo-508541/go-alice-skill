package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Visitor struct {
	ID      int
	Name    string
	Phones  []string
	Company string
}

var visitors = map[string]Visitor{
	"1": {
		ID:   1,
		Name: "Guest",
		Phones: []string{
			`789-673-56-90`,
			`612-934-77-23`,
		}},
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	var id string

	if req.Method == http.MethodPost {
		var visitor Visitor

		if err := json.NewDecoder(req.Body).Decode(&visitor); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// var buf bytes.Buffer
		// // читаем тело запроса
		// _, err := buf.ReadFrom(req.Body)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// // десериализуем JSON в Visitor
		// if err = json.Unmarshal(buf.Bytes(), &visitor); err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		id = strconv.Itoa(visitor.ID)
		// добавляем в мапу
		visitors[id] = visitor
	} else {
		id = req.URL.Query().Get("id")
	}
	resp, err := json.Marshal(visitors[id])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	go func() {
		time.Sleep(time.Second)
		http.Post(`http://localhost:8080`, `application/json`,
			// ключи указаны в разных регистрах, но данные сконвертируются правильно
			bytes.NewBufferString(`{"ID": 10, "NaMe": "Gopher", "company": "Don't Panic"}`))
	}()

	http.ListenAndServe("localhost:8080", http.HandlerFunc(JSONHandler))
}
