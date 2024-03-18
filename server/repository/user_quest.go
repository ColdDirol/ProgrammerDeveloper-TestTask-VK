package repository

import (
	"ProgrammerDeveloper-TestTask-VK/models"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"database/sql"
	"errors"
)

func AddUserQuest(userID int, questID int) error {
	var count int
	err := utils.DB.QueryRow(`
        SELECT COUNT(*) FROM users_quests WHERE user_id = $1 AND quest_id = $2
    `, userID, questID).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		return errors.New("quest already complete")
	}
	_, err = utils.DB.Exec(`
        INSERT INTO users_quests (user_id, quest_id) VALUES ($1, $2)
    `, userID, questID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]models.UserWithoutCredentials, error) {
	rows, err := utils.DB.Query(`
        SELECT id, first_name, last_name, balance, role FROM users
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserWithoutCredentials
	for rows.Next() {
		var user models.UserWithoutCredentials
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Balance, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(userID int) (*models.UserWithoutCredentials, error) {
	query := "SELECT id, first_name, last_name, balance, role FROM users WHERE id = $1"
	var user models.UserWithoutCredentials
	err := utils.DB.QueryRow(query, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Balance, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserWithCompletedQuestsByID(userID int) (*models.UserWithQuests, error) {
	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	quests, err := GetCompletedQuestsForUser(user.ID)
	if err != nil {
		return nil, err
	}
	userWithQuests := models.UserWithQuests{
		User:   *user,
		Quests: quests,
	}

	return &userWithQuests, nil
}

func GetAllUsersWithCompletedQuests() (*[]models.UserWithQuests, error) {
	users, err := GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersWithQuests []models.UserWithQuests

	for _, user := range users {
		quests, err := GetCompletedQuestsForUser(user.ID)
		if err != nil {
			return nil, err
		}
		userWithQuests := models.UserWithQuests{
			User:   user,
			Quests: quests,
		}
		usersWithQuests = append(usersWithQuests, userWithQuests)
	}

	return &usersWithQuests, nil
}

func GetCompletedQuestsForUser(userID int) ([]models.Quest, error) {
	rows, err := utils.DB.Query(`
        SELECT quests.id, quests.name, quests.cost
        FROM quests
        INNER JOIN users_quests ON quests.id = users_quests.quest_id
        WHERE users_quests.user_id = $1
    `, userID)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quests, nil
}

func UpdateUser(userID int, user models.UserWithoutCredentials) error {
	existingUser, err := GetUserByID(userID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}
	_, err = utils.DB.Exec(`
		UPDATE users 
		SET first_name = $1, last_name = $2, balance = $3, role = $4
		WHERE id = $5
	`, user.FirstName, user.LastName, user.Balance, user.Role, userID)
	if err != nil {
		return err
	}
	return nil
}

func CompleteQuest(userID int, questID int) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}
	quest, err := GetQuestByID(questID)
	if err != nil {
		return err
	}

	user.Balance += quest.Cost

	err = AddUserQuest(userID, questID)
	if err != nil {
		return err
	}
	err = UpdateUser(userID, *user)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAllUsersQuests() error {
	_, err := utils.DB.Exec(`
		DELETE FROM users_quests
	`)
	if err != nil {
		return errors.New("failed to clear users_actors table: " + err.Error())
	}
	return nil
}
