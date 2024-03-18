package handler

import (
	"ProgrammerDeveloper-TestTask-VK/auth/middleware"
	"ProgrammerDeveloper-TestTask-VK/server/service"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"net/http"
)

func InitUsersHandlers() {
	http.HandleFunc("/users", middleware.Middleware(usersHandler))
	http.HandleFunc("/users/quests/", middleware.Middleware(userQuestsHandler))
	http.HandleFunc("/users/complete_quest/", middleware.Middleware(completeQuestHandler))
	http.HandleFunc("/users/", middleware.Middleware(usersByIDHandler))
}

func userQuestsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractIDFromURL(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		service.GetUserWithQuestsByID(w, userID)
	} else {
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func completeQuestHandler(w http.ResponseWriter, r *http.Request) {
	userID, questID, err := utils.ExtractUserIDWithQuestID(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		service.CompleteQuest(w, userID, questID)
	} else {
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service.GetUsers(w)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func usersByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractIDFromURL(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		service.GetUserByID(w, userID)
	case http.MethodDelete:
		service.DeleteUser(w, userID)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}
