package storage

import (
	"database/sql"
	"financeTracker/internal/storage/queries"
	"financeTracker/pkg/config"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
	//TODO:здесь всторить разные репозитории по базе данных
}

func NewStorage(dbCfg *config.DatabaseConfig) (*Storage, error) {
	const op = `storage.NewStorage`

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbCfg.User, dbCfg.Password, dbCfg.DbName, dbCfg.SslMode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("op: %s, err: %w", op, err)
	}

	s := Storage{Db: db}
	err = s.CreateTables()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Storage) CreateTables() error {
	const op = `storage.CreateTables`

	tablesQueries := []string{
		queries.CreateUserQuery,
		queries.CreateGoalQuery,
		queries.CreateWalletQuery,
		queries.CreateSummaryQuery,
		queries.CreateCategoryQuery,
		queries.CreateExpenseQuery,
		queries.CreateIncomeQuery,
	}

	for i, query := range tablesQueries {
		_, err := s.Db.Exec(query)
		if err != nil {
			return fmt.Errorf("op: %s, query_index: %d, err: %w", op, i, err)
		}
	}
	return nil
}
