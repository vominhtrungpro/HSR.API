// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbmodel

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	Character       *character
	Element         *element
	Path            *path
	SchemaMigration *schemaMigration
	User            *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Character = &Q.Character
	Element = &Q.Element
	Path = &Q.Path
	SchemaMigration = &Q.SchemaMigration
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		Character:       newCharacter(db, opts...),
		Element:         newElement(db, opts...),
		Path:            newPath(db, opts...),
		SchemaMigration: newSchemaMigration(db, opts...),
		User:            newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Character       character
	Element         element
	Path            path
	SchemaMigration schemaMigration
	User            user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Character:       q.Character.clone(db),
		Element:         q.Element.clone(db),
		Path:            q.Path.clone(db),
		SchemaMigration: q.SchemaMigration.clone(db),
		User:            q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		Character:       q.Character.replaceDB(db),
		Element:         q.Element.replaceDB(db),
		Path:            q.Path.replaceDB(db),
		SchemaMigration: q.SchemaMigration.replaceDB(db),
		User:            q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Character       ICharacterDo
	Element         IElementDo
	Path            IPathDo
	SchemaMigration ISchemaMigrationDo
	User            IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Character:       q.Character.WithContext(ctx),
		Element:         q.Element.WithContext(ctx),
		Path:            q.Path.WithContext(ctx),
		SchemaMigration: q.SchemaMigration.WithContext(ctx),
		User:            q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
