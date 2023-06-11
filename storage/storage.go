package storage

import "github.com/MLaskun/go-workshop/types"

type Storage interface {
	GetClientByID(int) (*types.Client, error)
	CreateClient(*types.Client) error
	DeleteClient(int) error
}
