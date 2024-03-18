package main

import (
	"ProgrammerDeveloper-TestTask-VK/auth"
	"ProgrammerDeveloper-TestTask-VK/server/handler"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/database"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func waitForPostgres(database utils.Database) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", database.Host, database.Port, database.DBName, database.Username, database.Password)

	maxAttempts := 200
	attempt := 0
	for {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			utils.LOG.Error("Error connecting to PostgreSQL: %v", err)
		} else {
			defer db.Close()

			if err := db.Ping(); err == nil {
				utils.LOG.Info("PostgreSQL is up - executing command")
				break
			}
		}

		attempt++
		if attempt > maxAttempts {
			utils.LOG.Error("Exceeded max attempts, giving up.")
			os.Exit(1)
		}

		utils.LOG.Info("Wait for Postgres - waiting", slog.String("attempt", strconv.Itoa(attempt)), slog.String("attempt", strconv.Itoa(maxAttempts)))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	config := utils.InitConfig("config.json")

	log := logger.SetupLogger(config.Env)
	log.Info("You are in mode:", slog.String("env", config.Env))

	waitForPostgres(config.Database)

	db, err := database.InitDatabase(
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)
	if err != nil {
		log.Error("failed to init database", logger.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	auth.InitAuth(&config.Auth)
	handler.InitUsersHandlers()
	handler.InitQuestsHandlers()

	runApplication(&config.HTTPServer, log)
}

func runApplication(httpConfig *utils.HTTPServer, log *slog.Logger) {
	log.Info("Server start listening on port: ", slog.String("port", httpConfig.Port))
	httpAddress := httpConfig.Host + ":" + httpConfig.Port
	err := http.ListenAndServe(httpAddress, nil)
	if err != nil {
		log.Error("Failed to start server", logger.Err(err))
		os.Exit(1)
	}
}
