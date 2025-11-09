package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	usermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/user_model"
	"github.com/i-sub135/go-rest-blueprint/source/config"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/db"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
)

func main() {
	// Load config
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg := config.GetConfig()

	// Init logger
	logger.Init(cfg.Log.PrettyConsole)

	// Connect to database
	database, err := db.Init()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to database")
		log.Fatal(err)
	}

	logger.Info().Msg("Starting user migration...")

	// Auto migrate User model
	if err := database.AutoMigrate(&usermodel.User{}); err != nil {
		logger.Error().Err(err).Msg("Failed to migrate user table")
		log.Fatal(err)
	}

	logger.Info().Msg("User migration completed successfully!")

	// Generate 100 random users
	firstNames := []string{"John", "Jane", "Alex", "Sarah", "Mike", "Emma", "David", "Lisa", "Chris", "Anna", "Tom", "Maria", "James", "Linda", "Robert", "Patricia", "Michael", "Jennifer", "William", "Elizabeth"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin"}
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "example.com", "test.com", "company.com"}

	// Seed random
	rand.Seed(time.Now().UnixNano())

	logger.Info().Msg("Generating 100 random users...")

	var sampleUsers []usermodel.User
	for i := 1; i <= 100; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		domain := domains[rand.Intn(len(domains))]

		user := usermodel.User{
			Name: fmt.Sprintf("%s %s", firstName, lastName),
			Email: fmt.Sprintf("%s.%s%d@%s",
				firstName,
				lastName,
				rand.Intn(999)+1, // Random number 1-999
				domain),
		}
		sampleUsers = append(sampleUsers, user)
	}

	// Insert users
	successCount := 0
	for i, user := range sampleUsers {
		var existingUser usermodel.User
		if err := database.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			// User doesn't exist, create it
			if err := database.Create(&user).Error; err != nil {
				logger.Error().Err(err).Str("email", user.Email).Msg("Failed to create sample user")
			} else {
				successCount++
				logger.Info().Int("index", i+1).Str("name", user.Name).Str("email", user.Email).Msg("Sample user created")
			}
		} else {
			logger.Info().Str("email", user.Email).Msg("Sample user already exists")
		}
	}

	logger.Info().Int("total_created", successCount).Int("total_attempted", len(sampleUsers)).Msg("Migration process completed!")
}
