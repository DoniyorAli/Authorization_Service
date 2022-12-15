package storage

import (
	"UacademyGo/Blogpost/Authorization_Service/protogen/blogpost"
)

type StorageInter interface {
	//* Article
	AddNewArticle(id string, box *blogpost.CreateArticleRequest) error
	GetArticleById(id string) (*blogpost.GetArticleByIDResponse, error)
	GetArticleList(offset, limit int, search string) (dataset *blogpost.GetArticleListResponse, err error)
	UpdateArticle(box *blogpost.UpdateArticleRequest) error
	DeleteArticle(id string) error
	//* Author
	AddAuthor(id string, box *blogpost.CreateAuthorRequest) error
	GetAuthorById(id string) (*blogpost.Author, error)
	GetAuthorList(limit, offset int, search string) (dataset *blogpost.GetAuthorListResponse, err error)
	UpdateAuthor(box *blogpost.UpdateAuthorRequest) error
	DeleteAuthor(id string) error
}
