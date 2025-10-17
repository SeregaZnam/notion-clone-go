package env

import (
	"context"

	notionApi "github.com/SeregaZnam/notion-clone-go/internal/api/notion"
	database "github.com/SeregaZnam/notion-clone-go/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	Ctx           context.Context
	DB            *pgxpool.Pool
	NotionHandler *notionApi.Handler
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

	return nil
}

func CreateAndInit(ctx context.Context) (Env, error) {
	e := Env{
		Ctx: ctx,
	}

	err := e.populateDependencies()

	return e, err
}
