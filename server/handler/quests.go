package handler

import (
	"ProgrammerDeveloper-TestTask-VK/auth/middleware"
	"ProgrammerDeveloper-TestTask-VK/server/service"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"net/http"
)

func InitQuestsHandlers() {
	http.HandleFunc("/quests", middleware.Middleware(questsHandler))
	http.HandleFunc("/quests/", middleware.Middleware(questsByIDHandler))
}

func questsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service.GetQuests(w)
	case http.MethodPost:
		service.PostQuest(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func questsByIDHandler(w http.ResponseWriter, r *http.Request) {
	questID, err := utils.ExtractIDFromURL(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		service.GetQuestByID(w, questID)
	case http.MethodPut:
		service.UpdateQuest(w, r, questID)
	case http.MethodDelete:
		service.DeleteQuest(w, questID)
	default:
		utils.LOG.Error("invalid http method", http.StatusMethodNotAllowed)
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}
