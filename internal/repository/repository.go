package repository

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"sync"
)

type Repository interface {
	GetAddress(id int) (string, error)
	GetIDHist(request string) (int, error)
	GetAddressID(id int) (int, error)
	SaveSearchHist(request string) (int, error)
	SaveAddress(address string) (int, error)
	SaveHistSearchAddress(searchHistID, address int) error
}

type GeoRepo struct {
	db     *sql.DB
	sqlBlb sq.StatementBuilderType
	sync.Mutex
}

func New(db *sql.DB) GeoRepo {
	return GeoRepo{
		db:     db,
		sqlBlb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r GeoRepo) GetAddress(id int) (string, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `SELECT data FROM addresses WHERE id = $1`

	addr := ""
	if err := r.db.QueryRow(query, id).Scan(addr); err != nil {
		return addr, err
	}
	return addr, nil
}

func (r GeoRepo) GetIDHist(request string) (int, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `SELECT id FROM search_history WHERE similarity(query, $1) >= 0.7`
	id := 0
	if err := r.db.QueryRow(query, request).Scan(&id); err != nil {
		return id, nil
	}
	return id, nil
}

func (r GeoRepo) GetAddressID(id int) (int, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `SELECT address_id FROM history_search_address WHERE search_history_id = $1`
	addrId := 0
	if err := r.db.QueryRow(query, id).Scan(&addrId); err != nil {
		return addrId, err
	}
	return addrId, nil
}

func (r GeoRepo) SaveSearchHist(request string) (int, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `INSERT INTO search_history (query) VALUES ($1) RETURNING id`

	id := 0
	if err := r.db.QueryRow(query, request).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r GeoRepo) SaveAddress(address string) (int, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `INSERT INTO address (data) VALUES ($1) RETURNING id`
	id := 0
	if err := r.db.QueryRow(query, address).Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r GeoRepo) SaveHistSearchAddress(searchHistID, address int) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	query := `INSERT INTO history_search_address (search_history_id, address_id) VALUES ($1, $2)`

	if _, err := r.db.Exec(query, searchHistID, address); err != nil {
		return err
	}
	return nil
}
