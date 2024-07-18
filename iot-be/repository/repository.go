package repository

import (
	"context"
	"web-ivsr-be/model"
)

type Repo interface {
	GetDataRepo(ctx context.Context) ([]model.Information, error)
	PutDataRepo(ctx context.Context, nodeName string, value bool) error 
	GetDataIdRepo(ctx context.Context, node_id string) ([]model.Information, error)
}
