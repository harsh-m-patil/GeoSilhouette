package datatypes

type ScoreCard struct {
	minScore  int
	LivesLeft int
	Right     int
	Wrong     int
}

func NewScoreCard() *ScoreCard {
	return &ScoreCard{
		minScore:  0,
		LivesLeft: 5,
		Right:     0,
		Wrong:     0,
	}
}
