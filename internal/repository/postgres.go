package repository

import (
	"NoteKeeper/internal/domain"
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

const (
	ALIVE_TIMEOUT = 5 * time.Second
)

type Postgres struct {
	logger *zap.Logger
	db     *pg.DB
}

func NewPostgres(logger *zap.Logger) *Postgres {
	return &Postgres{
		logger: logger.Named("Postgres"),
	}
}

// TODO добавить кастомные ошибки
func (p *Postgres) Init(dsn string) error {
	connOpts, err := pg.ParseURL(dsn)
	if err != nil {
		p.logger.Error("parsing of Postgres DSN failed")
		return err
	}

	p.db = pg.Connect(connOpts)
	if _, err := p.Alive(); err != nil {
		p.logger.Error("Postgres is not alive")
		return err
	}

	err = p.db.Model((*domain.Note)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
		Temp:        false,
	})
	if err != nil {
		p.logger.Error("Creation of records table failed")
		return err
	}

	p.logger.Info("Storage successfully initialized")
	return nil
}

func (p *Postgres) Alive() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ALIVE_TIMEOUT)
	defer cancel()
	if err := p.db.Ping(ctx); err != nil {
		p.logger.Error("Postgres ping failed")
		return nil, err
	}
	return "alive", nil
}

func (p *Postgres) Insert(note domain.Note) error {
	if _, err := p.db.Model(&note).Insert(); err != nil {
		p.logger.Error("insert failed")
		return err
	}
	p.logger.Info("inserted successfully")
	return nil
}

func (p *Postgres) GetOne(uid uuid.UUID) (domain.Note, error) {
	note := domain.Note{
		UID: uid,
	}

	if err := p.db.Model(&note).WherePK().Select(); err != nil {
		p.logger.Error("getOne failed")
		return domain.Note{}, err
	}

	//if err := p.db.Model(&note).Where("uid = ?", uid).Select(); err != nil {
	//	p.logger.Error("getOne failed")
	//	return domain.Note{}, err
	//}
	p.logger.Info("got note successfully")
	return note, nil
}

func addOptions(q *pg.Query, opts domain.SearchOptions) {
	if opts.Title != nil {
		q = q.Where("title = ?", opts.Title)
	}

	if opts.Author != nil {
		q = q.Where("author = ?", opts.Author)
	}

	if opts.Tags != nil {
		q = q.Where("tags in (?)", pg.In(opts.Tags))
	}
}

func (p *Postgres) Get(opts domain.SearchOptions) ([]domain.Note, error) {
	var notes []domain.Note
	q := p.db.Model(&notes)
	addOptions(q, opts)

	if err := q.Select(); err != nil {
		p.logger.Error("get failed")
		return nil, err
	}
	p.logger.Info("got notes successfully")
	return notes, nil
}

func (p *Postgres) Update(note domain.Note) error {
	if _, err := p.db.Model(&note).WherePK().Update(); err != nil {
		p.logger.Error("update failed")
		return err
	}
	p.logger.Info("updated successfully")
	return nil
}

func (p *Postgres) Delete(uid uuid.UUID) error {
	if _, err := p.db.Model((*domain.Note)(nil)).Where("uid = ?", uid).Delete(); err != nil {
		p.logger.Error("delete failed")
		return err
	}
	p.logger.Info("deleted successfully")
	return nil
}
