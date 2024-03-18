package models

import "ProgrammerDeveloper-TestTask-VK/utils"

func CreateUserQuestTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users_quests (
            user_id INT,
            quest_id INT
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
