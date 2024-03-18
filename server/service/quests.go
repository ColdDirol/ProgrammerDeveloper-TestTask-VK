package service

import (
	"ProgrammerDeveloper-TestTask-VK/models"
	"ProgrammerDeveloper-TestTask-VK/server/repository"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"encoding/json"
	"net/http"
)

func GetQuests(w http.ResponseWriter) {
	quests, err := repository.GetAllQuests()
	if err != nil {
		utils.LOG.Error("failed to get quests", logger.Err(err))
		http.Error(w, "failed to get quests", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(quests)
	if err != nil {
		utils.LOG.Error("failed to marshal quests data", logger.Err(err))
		http.Error(w, "failed to marshal quests data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetQuestByID(w http.ResponseWriter, questID int) {
	quest, err := repository.GetQuestByID(questID)
	if err != nil {
		utils.LOG.Error("failed to get quest", logger.Err(err))
		http.Error(w, "failed to get quest", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(quest)
	if err != nil {
		utils.LOG.Error("failed to marshal quest data", logger.Err(err))
		http.Error(w, "failed to marshal quest data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func PostQuest(w http.ResponseWriter, r *http.Request) {
	var quest models.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)
	if err != nil {
		utils.LOG.Error("failed to decode quest data", logger.Err(err))
		http.Error(w, "failed to decode quest data", http.StatusInternalServerError)
		return
	}

	err = repository.AddQuest(quest)
	if err != nil {
		utils.LOG.Error("failed to insert quest", logger.Err(err))
		http.Error(w, "failed to insert quest", http.StatusBadRequest)
		return
	}
}

func UpdateQuest(w http.ResponseWriter, r *http.Request, questID int) {
	var quest models.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)
	if err != nil {
		utils.LOG.Error("failed to decode quest data", logger.Err(err))
		http.Error(w, "failed to decode quest data", http.StatusInternalServerError)
		return
	}

	err = repository.UpdateQuest(questID, quest)
	if err != nil {
		utils.LOG.Error("failed to update quest", logger.Err(err))
		http.Error(w, "failed to update quest", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteQuest(w http.ResponseWriter, questID int) {
	err := repository.DeleteQuestByID(questID)
	if err != nil {
		utils.LOG.Error("failed to delete quest", logger.Err(err))
		http.Error(w, "failed to delete quest", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
