package utils

import (
	"strconv"
	"strings"
)

func ExtractIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ExtractUserIDWithQuestID(path string) (int, int, error) {
	parts := strings.Split(path, "/")
	questIDStr := parts[len(parts)-1]
	questID, err := strconv.Atoi(questIDStr)
	if err != nil {
		return 0, 0, err
	}
	userIDStr := parts[len(parts)-2]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, 0, err
	}

	return userID, questID, nil
}
