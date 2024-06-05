package domain

// tipovi oklade
// “who wins the match”, “total number of goals”, “number of corners in 1. half”.
type Market struct {
	ID       string
	Name     string
	Status   int
	Outcomes []MarketOutcome
}

// rezultati koji se vezu za tip oklade
// “Team 1 will win the match”, “total number of goals will be 3”, “number of corners in 1. half will be 5”
type MarketOutcome struct {
	ID     string
	Name   string
	Status int
}

type MarketOutcomesPivot struct {
	MarketID  string
	OutcomeID string
}
