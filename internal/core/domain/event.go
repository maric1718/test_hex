package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// utakmica
// Describes a single event. Examples are “Real - Barcelona”, “BiH - Nigerija”...
type Event struct {
	ID       string
	Name     string    `json:"name"`
	StartsAt time.Time `json:"starts_at"`
	Status   int
	Markets  []Market
}

// “Who will win in event Real- Barcelona”, “Will there be an offside in event BiH - Nigerija”...
type EventMarket struct {
	ID     string
	Market Market
	Status int // event market status - market moze biti inactive samo za ovaj event
}

// rezultati koji se odnose direktno na event
type EventMarketOutcome struct {
	ID      string
	Outcome MarketOutcome
	Status  int
	Odd     float32
}

// TEMP
func LoadInitEvents() ([]Event, error) {
	var events []Event

	path := "internal/adapter/storage/file_system/json_data/initial/events.json"

	content, err := os.ReadFile(path)
	if err != nil {
		return events, err
	}

	if err := json.Unmarshal(content, &events); err != nil {
		return events, err
	}

	return events, nil
}

func (e *Event) UnmarshalJSON(data []byte) (err error) {
	type Alias Event

	aux := &struct {
		StartsAt string `json:"starts_at"`
		Name     string `json:"name"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.StartsAt != "" {
		// s, err := time.Parse("2006-01-02T15:04:05.000Z", aux.StartsAt)
		s, err := time.Parse(time.RFC3339, aux.StartsAt)
		if err != nil {
			fmt.Println("error_parsing") // TEMP
			return err
		}

		e.StartsAt = s
	}

	return nil
}

func LoadInitMarkets() ([]Market, error) {
	var markets []Market

	path := "internal/adapter/storage/file_system/json_data/initial/markets.json"

	content, err := os.ReadFile(path)
	if err != nil {
		return markets, err
	}

	if err := json.Unmarshal(content, &markets); err != nil {
		return markets, err
	}

	return markets, nil
}
