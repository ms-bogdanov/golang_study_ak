package main

import "encoding/json"

type User struct {
	Name     string
	Age      int
	Comments []Comment
}

type Comment struct {
	Text string
}

func getJSON(data []User) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
