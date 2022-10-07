package models

import "github.com/google/uuid"

type Key struct {
	Success bool
	Record  struct {
		Key              uuid.UUID // API key
		Owner            uuid.UUID // UUID of owner of API key
		Limit            int       // Limit of total queries
		QueriesInPastMin int       // Queries in the past minute
		TotalQueries     int       // Total count of queries
	}
}
