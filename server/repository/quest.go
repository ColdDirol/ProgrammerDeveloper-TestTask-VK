package repository

import (
	"ProgrammerDeveloper-TestTask-VK/models"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"database/sql"
	"errors"
)

func AddQuest(quest models.Quest) error {
	_, err := utils.DB.Exec(`
		INSERT INTO quests (name, cost) VALUES ($1, $2)
	`, quest.Name, quest.Cost)
	if err != nil {
		return err
	}
	return nil
}

func GetAllQuests() ([]models.Quest, error) {
	rows, err := utils.DB.Query(`
		SELECT id, name, cost FROM quests
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quests []models.Quest
	for rows.Next() {
		var quest models.Quest
		err := rows.Scan(&quest.ID, &quest.Name, &quest.Cost)
		if err != nil {
			return nil, err
		}
		quests = append(quests, quest)
	}
	return quests, nil
}

func GetQuestByID(questID int) (models.Quest, error) {
	var quest models.Quest
	err := utils.DB.QueryRow(`
		SELECT id, name, cost FROM quests WHERE id=$1
	`, questID).Scan(&quest.ID, &quest.Name, &quest.Cost)
	if err != nil {
		if err == sql.ErrNoRows {
			return quest, sql.ErrNoRows
		}
		return quest, err
	}
	return quest, nil
}

func UpdateQuest(questID int, newQuest models.Quest) error {
	_, err := utils.DB.Exec(`
		UPDATE quests SET name=$1, cost=$2 WHERE id=$3
	`, newQuest.Name, newQuest.Cost, questID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteQuestByID(id int) error {
	_, err := utils.DB.Exec(`
		DELETE FROM quests 
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllQuests() error {
	_, err := utils.DB.Exec(`
		DELETE FROM quests
	`)
	if err != nil {
		return errors.New("failed to delete actors: " + err.Error())
	}
	DeleteAllUsersQuests()
	return nil
}
