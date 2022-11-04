package spine

import (
	"context"
	goErr "errors"
	"github.com/google/uuid"
	"go-rest-webserver-template/pkg/errors"
	"go-rest-webserver-template/pkg/spine/db"

	"github.com/jinzhu/gorm"
)

type Repo struct {
	Db *db.DB
}

// FindByID fetches the record which matches the ID provided from the entity defined by receiver
func (repo Repo) FindByID(ctx context.Context, receiver IModel, id uuid.UUID) errors.IError {

	q := repo.Db.Instance(ctx).Where("id = ?", id).First(receiver)

	return GetDBError(q)
}

// FindByColumn fetches the record which matches the culumn name and Value provided from the entity defined by receiver
func (repo Repo) FindByColumn(ctx context.Context, receiver IModel, key string, val string) errors.IError {

	q := repo.Db.Instance(ctx).Where(key+"= ?", val).Order("created_at DESC").First(receiver)

	return GetDBError(q)
}

// FindByIDs fetches the all the records which matches the IDs provided from the entity defined by receivers
func (repo Repo) FindByIDs(ctx context.Context, receivers interface{}, ids []uuid.UUID) errors.IError {

	q := repo.Db.Instance(ctx).Where(AttributeID+" in (?)", ids).Find(receivers)

	return GetDBError(q)
}

// Create inserts a new record in the entity defined by the receiver
func (repo Repo) Create(ctx context.Context, receiver IModel) errors.IError {

	if err := receiver.SetDefaults(); err != nil {
		return err
	}

	if err := receiver.Validate(); err != nil {
		return err
	}

	q := repo.Db.Instance(ctx).Create(receiver)

	return GetDBError(q)
}

func (repo Repo) IsTransactionActive(ctx context.Context) bool {
	_, ok := ctx.Value(db.ContextKeyDatabase).(*gorm.DB)
	return ok
}

// Update will update the given model with respect to primary key / id available in it.
func (repo Repo) Update(ctx context.Context, receiver IModel, selectiveList ...string) errors.IError {

	q := repo.Db.Instance(ctx).Model(receiver)

	if len(selectiveList) > 0 {
		q = q.Select(selectiveList)
	}

	q = q.Update(receiver)

	if q.RowsAffected == 0 {
		return NoRowAffected.
			New(errNoRowAffected).
			Wrap(goErr.New("no rows have been updated"))
	}

	return GetDBError(q)
}

// Delete deletes the given model
func (repo Repo) Delete(ctx context.Context, receiver IModel) errors.IError {

	q := repo.Db.Instance(ctx).Delete(receiver)

	return GetDBError(q)
}

func (repo Repo) FindMany(
	ctx context.Context,
	receivers interface{},
	condition map[string]interface{}) errors.IError {

	q := repo.Db.Instance(ctx).Where(condition).Find(receivers)

	return GetDBError(q)
}

// Transaction will manage the execution inside a transactions
func (repo Repo) Transaction(ctx context.Context, fc func(ctx context.Context) errors.IError) (err errors.IError) {
	// If there is an active transaction then do not create a new transactions
	if _, ok := ctx.Value(db.ContextKeyDatabase).(*gorm.DB); ok {
		return fc(ctx)
	}

	panicked := true
	tx := repo.Db.Instance(ctx).Begin()

	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	// call the transaction handled with tx added in the context key
	err = fc(context.WithValue(ctx, db.ContextKeyDatabase, tx))
	if err == nil {
		err = GetDBError(tx.Commit())
	}
	panicked = false
	return
}
