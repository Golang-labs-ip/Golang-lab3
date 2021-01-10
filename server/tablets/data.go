package tablets

import (
	"time"
	"database/sql"
	"fmt"
)

type FullList struct {
	tabletID	int64 `json:"tabletID"`
	telemetryID	int64 `json:"tabletID"`
}

type Tablet struct {
	id		int64    `json:"id"`
	name	string `json:"name"`
}

type Telemetry struct {
	id				int64 `json:"id"`
	battery			int64 `json:"battery"`
	deviceTime		string `json:"name"`
	timeStamp		string `json:"telemetryId"`
	currentVideo	string `json:"currentVideo"`
}

type Tablets struct {
	TabletsArr []*Tablet `json:"Tablets"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListTablets() (*Tablets, error) {}

func (s *Store) Telemetry(id int64) ([]string, error) {}

func (s *Store) UpdateTablet(id int64) error {
	t := time.Now()
	time := t.Format("2006-06-01T14:15:16.123Z") 
	fmt.Println(time);
	_, err := s.Db.Exec("update VirtualMachine set isWorking=$1 where id=$2", time, id)
	return err
}


