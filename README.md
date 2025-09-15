POST-запрос

```
data := `{"name": "Иванов Иван", "email": "ivan@example.com"}`
resp, err := http.Post("https://example.com/adduser", "application/json", strings.NewReader(data))
if err != nil {
    return err
}
```
