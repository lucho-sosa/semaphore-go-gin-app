package main

import (
  "errors"
)

type article struct {
  ID int `json:"id"`
  Title   string `json:"title"`
  Content string `json:"content"`
}

var articleList = []article{
  article{ID: 1, Title: "Hello", Content: "Article Content"},
  article{ID: 2, Title: "Hello 2", Content: "Article Content 2"},
}

func getAllArticles() []article {
  return articleList
}

func getArticleByID(id int) (*article, error) {
  for _, a := range articleList {
    if a.ID == id {
      return &a, nil
    }
  }
  return nil, errors.New("Article not found")
}
