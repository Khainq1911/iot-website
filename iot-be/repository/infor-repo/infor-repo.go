package inforrepo

import (
	"context"
	"fmt"
	"web-ivsr-be/database"
	"web-ivsr-be/model"
	"web-ivsr-be/repository"
)

type RepoImpl struct {
	Sql *database.Sql
}

func NewRepo(sql *database.Sql) repository.Repo {
	return (&RepoImpl{
		Sql: sql,
	})
}

func (db *RepoImpl) GetDataRepo(ctx context.Context) ([]model.Information, error) {
	InforModel := []model.Information{}

	query := `select * from datatable`

	if err := db.Sql.Db.Select(&InforModel, query); err != nil {
		return nil, err
	}

	return InforModel, nil
}

func (db *RepoImpl) GetDataIdRepo(ctx context.Context, node_id string) ([]model.Information, error){
	InforModel := []model.Information{}

	query := `select * from datatable where node_id = $1`
	if err := db.Sql.Db.Get(&InforModel, query, node_id); err != nil {
		return InforModel, err
	}

	return InforModel, nil
}
func (db *RepoImpl) PutDataRepo(ctx context.Context, nodeName string, value bool) error {
	query := `update datatable set status = $1 where node_id = $2`
	_, err := db.Sql.Db.Exec(query, value, nodeName)
	if err != nil{
		fmt.Println("error o put repo")
		return err
	}
	return nil
}