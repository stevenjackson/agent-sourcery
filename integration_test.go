package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: 2024/07/18:
// Write a test that determines the current assignment
// Clean everything up

func TestGiveMeAllOfDalesEngagements(t *testing.T) {
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/agent_sourcery_test"

	db, _ := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	db.AutoMigrate(&Event{})

	t.Run("Can get list of engagements for an agent", func(t *testing.T) {

		insertEvent(db, "agent-00116", "AgentAdded", AgentAdded{
			AgentName: "Dale Karp",
			StartDate: "2021-09-30",
		})

		insertEvent(db, "engagement-123456789", "EngagementStarted", EngagementStarted{
			StartDate:  "2021-09-30",
			ClientName: "Handshake",
		})
		insertEvent(db, "agent-00116", "AssignmentStarted", AssignmentStarted{
			StartDate:  "2021-09-30",
			ClientName: "Handshake",
			AgentName:  "Dale Karp",
		})
		insertEvent(db, "agent-00116", "AssignmentEnded", AssignmentEnded{
			EndDate:    "2021-10-30",
			ClientName: "Handshake",
			AgentName:  "Dale Karp",
		})
		insertEvent(db, "engagement-234567890", "EngagementStarted", EngagementStarted{
			StartDate:  "2021-10-31",
			ClientName: "CARS",
		})
		insertEvent(db, "agent-00116", "AssignmentStarted", AssignmentStarted{
			StartDate:  "2021-09-30",
			ClientName: "CARS",
			AgentName:  "Dale Karp",
		})

		clients := ClientHistory(db, "Dale Karp")
		assert.Equal(t, []string{"Handshake", "CARS"}, clients)
	})

	db.Exec("DELETE FROM events")
}

func insertEvent(db *gorm.DB, streamName string, eventType string, data interface{}) Event {
	d, _ := json.Marshal(data)
	event := Event{
		StreamName: streamName,
		EventType:  eventType,
		Data:       string(d),
	}

	db.Save(&event)
	return event
}
