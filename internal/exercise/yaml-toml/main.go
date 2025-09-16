package main

import (
	"fmt"

	toml "github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	var v Data
	err := yaml.Unmarshal([]byte(yamlData), &v)
	if err != nil {
		panic(err)
	}
	out, err := toml.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
