package env

import (
	"context"

	notionApi "github.com/SeregaZnam/notion-clone-go/internal/api/handlers/notion"
	pageApi "github.com/SeregaZnam/notion-clone-go/internal/api/handlers/page"
	database "github.com/SeregaZnam/notion-clone-go/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	Ctx           context.Context
	DB            *pgxpool.Pool
	NotionHandler *notionApi.Handler
	PageHandler   *pageApi.Handler
}

func (e *Env) populateDependencies() error {
	// init DB
	db, err := database.InitDB(e.Ctx)
	if err != nil {
		return err
	}
	e.DB = db

	// init handlers with deps
	e.NotionHandler = notionApi.NewRepository(db)
	e.PageHandler = pageApi.NewRepository(db)

	return nil
}

func CreateAndInit(ctx context.Context) (Env, error) {
	e := Env{
		Ctx: ctx,
	}

	err := e.populateDependencies()

	return e, err
}
