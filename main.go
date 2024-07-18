package main

import (
	"encoding/json"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EngagementProposed struct {
	ClientName         string
	RequiredSkills     []string
	AgentCount         int
	ClientServicesLead string
}

type EngagementStarted struct {
	StartDate  string
	ClientName string
}

type EngagementEnded struct {
	EndDate    string
	ClientName string
}

type FTOUsed struct {
	StartDate string
	EndDate   string
	AgentName string
}

type AssignmentStarted struct {
	AgentName  string
	ClientName string
	StartDate  string
}

type AssignmentEnded struct {
	AgentName  string
	ClientName string
	EndDate    string
}

type SkillAdded struct {
	AgentName string
	SkillName string
	Level     int
}

type SkillLevelChange struct {
	AgentName string
	SkillName string
	Level     int
}

type InterestAdded struct {
	AgentName    string
	InterestName string
}

type InterestRemoved struct {
	AgentName    string
	InterestName string
}

type AgentAdded struct {
	AgentName string
	StartDate string
}

type AgentRemoved struct {
	AgentName  string
	RemoveDate string
}

type Event struct {
	gorm.Model
	StreamName string
	EventType  string
	Data       string
}

func main() {
	// TODO: 2024/06/27:
	// Write a test that adds an agent/engagement & query the state of the DB
	// e.g.: Give me all of Dale's engagmenets to date
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Event{})
}

// Projections
type Engagement struct {
}

func ClientHistory(db *gorm.DB, agentName string) []string {
	var events []Event
	db.Find(&events)
	var clientNames []string

	for _, event := range events {
		switch event.EventType {
		case "AssignmentStarted":
			var as AssignmentStarted
			json.Unmarshal([]byte(event.Data), &as)
			clientNames = append(clientNames, as.ClientName)
		}
	}
	return clientNames
}
