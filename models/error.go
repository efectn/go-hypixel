package models

type Error struct {
	Success bool   // Situation of process
	Cause   string // Cause of JSON error
}
