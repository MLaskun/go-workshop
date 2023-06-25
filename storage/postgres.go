package storage

import (
	"database/sql"
	"fmt"

	"github.com/MLaskun/go-workshop/types"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=goworkshop sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createClientAndVehicleTable()
}

func (s *PostgresStorage) createClientAndVehicleTable() error {
	query := `create table if not exists client(
		client_id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		email varchar(50),
		phone_number serial
	);
	create table if not exists vehicle(
		vehicle_id serial primary key,
		owner_id serial references client(client_id),
		vin varchar(50),
		vehicle_type varchar(50),
		make varchar(50),
		model varchar(50),
		year_of_production int
	)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateClient(cli *types.Client) error {
	query := `insert into client
	(first_name, last_name, email, phone_number)
	values($1, $2, $3, $4)`
	_, err := s.db.Query(
		query,
		cli.FirstName,
		cli.LastName,
		cli.Email,
		cli.PhoneNO)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetClients() ([]*types.Client, error) {
	rows, err := s.db.Query(`select * from client`)
	if err != nil {
		return nil, err
	}

	clients := []*types.Client{}
	for rows.Next() {
		client, err := scanIntoClient(rows)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (s *PostgresStorage) GetClientByID(id int) (*types.Client, error) {
	rows, err := s.db.Query(`select * from client where client_id = $1`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		client := &types.Client{}
		err := rows.Scan(
			&client.ID,
			&client.FirstName,
			&client.LastName,
			&client.Email,
			&client.PhoneNO)

		return client, err
	}

	return nil, fmt.Errorf("client with id [%d] not found", id)
}

func (s *PostgresStorage) DeleteClient(id int) error {
	_, err := s.db.Query(`delete from client where client_id = $1`, id)
	return err
}

func scanIntoClient(rows *sql.Rows) (*types.Client, error) {
	client := &types.Client{}
	err := rows.Scan(
		&client.ID,
		&client.FirstName,
		&client.LastName,
		&client.Email,
		&client.PhoneNO)

	return client, err
}
