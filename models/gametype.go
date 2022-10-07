package models

type GameType struct {
	ID           uint   // IDs are used when the API references a specific GameType.
	DatabaseName string // Database names are used when the API references a specific GameType.
	CleanName    string // Clean names are that is displayed to the user when referencing the name.
}
