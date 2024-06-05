CREATE TABLE "market_outcomes_pivot" (
    "market_id" integer NOT NULL,
    "outcome_id" integer NOT NULL,
    FOREIGN KEY (market_id) REFERENCES markets(id),
    FOREIGN KEY (outcome_id) REFERENCES market_outcomes(id),
    PRIMARY KEY ("market_id", "outcome_id")
);

