package types

type ScoreCard struct {
	minScore     int
	initScore    int
	CurrentScore int
}

func NewScoreCard() *ScoreCard {
	return &ScoreCard{
		minScore:     0,
		initScore:    50,
		CurrentScore: 50,
	}
}
