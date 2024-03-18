package models

import "ProgrammerDeveloper-TestTask-VK/utils"

type Quest struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Cost int    `json:"cost" db:"cost"`
}

func CreateQuestTable() error {
	_, err := utils.DB.Exec(`
		CREATE TABLE IF NOT EXISTS quests (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR(100),
		    cost INTEGER
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
