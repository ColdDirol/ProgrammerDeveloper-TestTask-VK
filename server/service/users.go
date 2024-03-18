package service

import (
	repository2 "ProgrammerDeveloper-TestTask-VK/auth/repository"
	"ProgrammerDeveloper-TestTask-VK/server/repository"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter) {
	users, err := repository.GetAllUsers()
	if err != nil {
		utils.LOG.Error("failed to get users", logger.Err(err))
		http.Error(w, "failed to get users", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		utils.LOG.Error("failed to marshal users data", logger.Err(err))
		http.Error(w, "failed to marshal users data", http.StatusInternalServerError)
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

func GetUserWithQuestsByID(w http.ResponseWriter, userID int) {
	user, err := repository.GetUserWithCompletedQuestsByID(userID)
	if err != nil {
		utils.LOG.Error("failed to get user", logger.Err(err))
		http.Error(w, "failed to get user", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		utils.LOG.Error("failed to marshal user data", logger.Err(err))
		http.Error(w, "failed to marshal user data", http.StatusInternalServerError)
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

func GetUsersWithQuests(w http.ResponseWriter) {
	users, err := repository.GetAllUsersWithCompletedQuests()
	if err != nil {
		utils.LOG.Error("failed to get users", logger.Err(err))
		http.Error(w, "failed to get users", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		utils.LOG.Error("failed to marshal users data", logger.Err(err))
		http.Error(w, "failed to marshal users data", http.StatusInternalServerError)
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

func GetUserByID(w http.ResponseWriter, userID int) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		utils.LOG.Error("failed to get user", logger.Err(err))
		http.Error(w, "failed to get user", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		utils.LOG.Error("failed to marshal user data", logger.Err(err))
		http.Error(w, "failed to marshal user data", http.StatusInternalServerError)
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

func DeleteUser(w http.ResponseWriter, userID int) {
	err := repository2.DeleteUserByID(userID)
	if err != nil {
		utils.LOG.Error("failed to delete user", logger.Err(err))
		http.Error(w, "failed to delete user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CompleteQuest(w http.ResponseWriter, userID int, questID int) {
	err := repository.CompleteQuest(userID, questID)
	if err != nil {
		utils.LOG.Error("failed to complete quest for this user", logger.Err(err), slog.String("userID", strconv.Itoa(userID)), slog.String("questID", strconv.Itoa(questID)))
		http.Error(w, "failed to complete quest for this user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
