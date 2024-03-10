package repository

type SearchHistory struct {
	ID    int    `db:"id" db_type:"SERIAL PRIMARY KEY"`
	Query string `db:"query" db_type:"VARCHAR(255)"`
}

func (h SearchHistory) TableName() string {
	return "search_history"
}

type Address struct {
	ID   int    `db:"id" db_type:"SERIAL PRIMARY KEY" json:"id"`
	Data string `db:"data" db_type:"VARCHAR(255)" json:"data"`
}

func (a Address) TableName() string {
	return "address"
}

type HistorySearchAddress struct {
	SearchHistoryID int `db:"search_history_id" db_type:"INTEGER"`
	AddressID       int `db:"address_id" db_type:"INTEGER"`
}

func (h HistorySearchAddress) TableName() string {
	return "history_search_address"
}
