package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	customermodel "github.com/i-sub135/go-rest-blueprint/source/common/model/customer_model"
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

	logger.Info().Msg("Starting customer migration...")

	// Auto migrate Customer model
	if err := database.AutoMigrate(&customermodel.Customer{}); err != nil {
		logger.Error().Err(err).Msg("Failed to migrate customer table")
		log.Fatal(err)
	}

	logger.Info().Msg("Customer migration completed successfully!")

	// Generate sample customer data
	firstNames := []string{"John", "Jane", "Alex", "Sarah", "Mike", "Emma", "David", "Lisa", "Chris", "Anna", "Tom", "Maria", "James", "Linda", "Robert", "Patricia", "Michael", "Jennifer", "William", "Elizabeth"}
	lastNames := []string{"Wijaya", "Santoso", "Kurniawan", "Sari", "Pratama", "Utomo", "Handayani", "Susanto", "Maharani", "Gunawan", "Fitria", "Permana", "Rahayu", "Nugroho", "Safitri", "Hidayat", "Wulandari", "Setiawan", "Anggraini", "Putra"}
	cities := []string{"Jakarta", "Surabaya", "Bandung", "Medan", "Semarang", "Makassar", "Palembang", "Tangerang", "Depok", "Bekasi", "Solo", "Batam", "Pekanbaru", "Bandar Lampung", "Malang", "Yogyakarta", "Bogor", "Denpasar", "Samarinda", "Balikpapan"}
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "company.id", "email.com"}
	addresses := []string{"Jl. Sudirman", "Jl. Thamrin", "Jl. Gatot Subroto", "Jl. Kuningan", "Jl. Senayan", "Jl. Kemang", "Jl. Pondok Indah", "Jl. Kelapa Gading", "Jl. Pluit", "Jl. PIK"}

	// Seed random
	rand.Seed(time.Now().UnixNano())

	logger.Info().Msg("Generating 50 sample customers...")

	var sampleCustomers []customermodel.Customer
	for i := 1; i <= 50; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		city := cities[rand.Intn(len(cities))]
		domain := domains[rand.Intn(len(domains))]
		address := addresses[rand.Intn(len(addresses))]

		// Generate random date of birth (age 18-65)
		minAge := 18 * 365 * 24 * time.Hour
		maxAge := 65 * 365 * 24 * time.Hour
		randomAge := minAge + time.Duration(rand.Int63n(int64(maxAge-minAge)))
		dateOfBirth := time.Now().Add(-randomAge)

		// Generate phone number
		phone := fmt.Sprintf("+628%d%d",
			rand.Intn(9)+1,
			rand.Intn(90000000)+10000000)

		customer := customermodel.Customer{
			FirstName: firstName,
			LastName:  lastName,
			Email: fmt.Sprintf("%s.%s%d@%s",
				firstName,
				lastName,
				rand.Intn(999)+1,
				domain),
			Phone:       phone,
			Address:     fmt.Sprintf("%s No. %d", address, rand.Intn(100)+1),
			City:        city,
			Country:     "Indonesia",
			DateOfBirth: &dateOfBirth,
			IsActive:    rand.Float32() > 0.1, // 90% active
		}
		sampleCustomers = append(sampleCustomers, customer)
	}

	// Insert customers
	successCount := 0
	for i, customer := range sampleCustomers {
		var existingCustomer customermodel.Customer
		if err := database.Where("email = ?", customer.Email).First(&existingCustomer).Error; err != nil {
			// Customer doesn't exist, create it
			if err := database.Create(&customer).Error; err != nil {
				logger.Error().Err(err).Str("email", customer.Email).Msg("Failed to create sample customer")
			} else {
				successCount++
				logger.Info().
					Int("index", i+1).
					Str("name", customer.FullName()).
					Str("email", customer.Email).
					Str("city", customer.City).
					Str("phone", customer.Phone).
					Bool("is_active", customer.IsActive).
					Msg("Sample customer created")
			}
		} else {
			logger.Info().Str("email", customer.Email).Msg("Sample customer already exists")
		}
	}

	logger.Info().
		Int("total_created", successCount).
		Int("total_attempted", len(sampleCustomers)).
		Msg("Customer migration process completed!")
}
