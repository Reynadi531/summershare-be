package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

func GeneratePostId() string {
	id, _ := gonanoid.New(10)
	return id
}
