package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/hs-zavet/news-radar/internal/app"
	"github.com/hs-zavet/news-radar/internal/app/models"
	"github.com/hs-zavet/news-radar/internal/config"
	"github.com/hs-zavet/news-radar/internal/content"
	"github.com/sirupsen/logrus"
)

type App interface {
	CreateArticle(ctx context.Context, request app.CreateArticleRequest) (models.Article, error)
	GetArticleByID(ctx context.Context, articleID uuid.UUID) (models.Article, error)
	UpdateArticle(ctx context.Context, articleID uuid.UUID, request app.UpdateArticleRequest) (models.Article, error)
	DeleteArticle(ctx context.Context, articleID uuid.UUID) error

	SetArticleTags(ctx context.Context, articleID uuid.UUID, tags []string) error
	GetArticleTags(ctx context.Context, articleID uuid.UUID) ([]models.Tag, error)
	AddArticleTag(ctx context.Context, articleID uuid.UUID, tag string) error
	DeleteArticleTag(ctx context.Context, articleID uuid.UUID, tag string) error
	CleanArticleTags(ctx context.Context, articleID uuid.UUID) error

	SetAuthors(ctx context.Context, articleID uuid.UUID, authors []uuid.UUID) error
	GetArticleForAuthors(ctx context.Context, articleID uuid.UUID) ([]models.Article, error)
	AddArticleAuthor(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error
	DeleteArticleAuthor(ctx context.Context, articleID uuid.UUID, authorID uuid.UUID) error
	CleanArticleAuthors(ctx context.Context, articleID uuid.UUID) error

	UpdateArticleContent(ctx context.Context, articleID uuid.UUID, index int, section content.Section) error

	CreateTag(ctx context.Context, request app.CreateTagRequest) error
	DeleteTag(ctx context.Context, name string) error
	UpdateTag(ctx context.Context, name string, request app.UpdateTagRequest) error
	GetTag(ctx context.Context, name string) (models.Tag, error)

	CreateAuthor(ctx context.Context, request app.CreateAuthorRequest) error
	UpdateAuthor(ctx context.Context, authorID uuid.UUID, request app.UpdateAuthorRequest) error
	DeleteAuthor(ctx context.Context, authorID uuid.UUID) error
	GetAuthorByID(ctx context.Context, authorID uuid.UUID) (models.Author, error)
}

type Handler struct {
	app      App
	cfg      config.Config
	log      *logrus.Entry
	upgrader websocket.Upgrader
}

func NewHandlers(cfg config.Config, log *logrus.Entry, app *app.App) Handler {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			//allowedOrigin := "https://example.com"
			//return r.Header.GetTag("Origin") == allowedOrigin
			return true
		},
	}

	return Handler{
		app:      app,
		cfg:      cfg,
		log:      log,
		upgrader: upgrader,
	}
}
