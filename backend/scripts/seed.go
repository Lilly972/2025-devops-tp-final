package main

import (
	"context"
	"log"

	"github.com/anthonyquere/christmas-gift-list/internal/database"
	"github.com/anthonyquere/christmas-gift-list/internal/models"
	"github.com/anthonyquere/christmas-gift-list/internal/repository"
)

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	ctx := context.Background()
	personRepo := repository.NewPersonRepository()
	giftRepo := repository.NewGiftRepository()

	log.Println("Seeding database...")

	// Create people
	people := []models.CreatePersonInput{
		{Name: "Santa Claus"},
		{Name: "Mrs. Claus"},
		{Name: "Rudolph"},
		{Name: "Frosty the Snowman"},
	}

	createdPeople := make([]*models.Person, 0)
	for _, p := range people {
		person, err := personRepo.Create(ctx, p)
		if err != nil {
			log.Printf("Failed to create person %s: %v", p.Name, err)
			continue
		}
		createdPeople = append(createdPeople, person)
		log.Printf("Created person: %s", person.Name)
	}

	// Create gifts for each person
	giftIdeas := map[string][]models.CreateGiftInput{
		"Santa Claus": {
			{Title: "New Sleigh", Description: "A modern, eco-friendly sleigh with GPS navigation"},
			{Title: "Comfortable Boots", Description: "Warm, waterproof boots for chimney climbing"},
			{Title: "Cookie Recipe Book", Description: "From around the world"},
		},
		"Mrs. Claus": {
			{Title: "Baking Equipment", Description: "Professional cookie cutters and mixers"},
			{Title: "Knitting Supplies", Description: "Finest wool and needles for sweater making"},
			{Title: "Spa Day", Description: "A relaxing day after busy Christmas season"},
		},
		"Rudolph": {
			{Title: "Nose Polish", Description: "Extra bright formula for foggy nights"},
			{Title: "Carrot Subscription", Description: "Premium organic carrots delivered monthly"},
			{Title: "Reindeer Games Trophy", Description: "For being the MVP"},
		},
		"Frosty the Snowman": {
			{Title: "Top Hat Upgrade", Description: "Magical silk top hat with built-in cooling"},
			{Title: "Carrot Nose Collection", Description: "Different colors and sizes"},
			{Title: "Winter Coat", Description: "Extra insulation for warmer days"},
		},
	}

	for _, person := range createdPeople {
		gifts, ok := giftIdeas[person.Name]
		if !ok {
			continue
		}

		for _, giftInput := range gifts {
			giftInput.PersonID = person.ID
			gift, err := giftRepo.Create(ctx, giftInput)
			if err != nil {
				log.Printf("Failed to create gift for %s: %v", person.Name, err)
				continue
			}
			log.Printf("  Created gift: %s for %s", gift.Title, person.Name)
		}
	}

	log.Println("Database seeding completed!")
}
