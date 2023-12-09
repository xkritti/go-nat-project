package models

type MathAnalytic struct {
	Classification string           `json:"classification"`
	Desc           string           `json:"desc"`
	Parts          MathPartAnalytic `json:"parts"`
}

type MathPartAnalytic struct {
	Calculation     string `json:"calculation"`
	ProblemSolution string `json:"problem_solution"`
	Appliation      string `json:"appliation"`
}

type SciAnalytic struct {
	Classification string          `json:"classification"`
	Desc           string          `json:"desc"`
	Parts          SciPartAnalytic `json:"parts"`
}

type SciPartAnalytic struct {
	Lesson     string `json:"lesson"`
	Appliation string `json:"appliation"`
}

type EngAnalytic struct {
	Classification string          `json:"classification"`
	Desc           string          `json:"desc"`
	Parts          EngPartAnalytic `json:"parts"`
}

type EngPartAnalytic struct {
	Expression string `json:"expression"`
	Reading    string `json:"reading"`
	Structure  string `json:"structure"`
	Vocabulary string `json:"vocabulary"`
}

type GetMathAnalyticRequest struct {
	ScorePercentage  float32 `json:"score_percentage" validate:"min=0.00,max=100.00"`
	CalPartScore     float32 `json:"cal_part_score" validate:"min=0.00,max=22.6"`
	ProblemPartScore float32 `json:"problem_part_score" validate:"min=0.00,max=56.65"`
	AppliedPartScore float32 `json:"applied_part_score" validate:"min=0.00,max=24.75"`
}

type GetSciAnalyticRequest struct {
	ScorePercentage  float32 `json:"score_percentage" validate:"min=0.00,max=100.00"`
	LessonPartScore  float32 `json:"lesson_part_score" validate:"min=0.00,max=80.00"`
	AppliedPartScore float32 `json:"applied_part_score" validate:"min=0.00,max=19.50"`
}

type GetEngAnalyticRequest struct {
	ScorePercentage     float32 `json:"score_percentage" validate:"min=0.00,max=100.00"`
	ExpressionPartScore float32 `json:"exp_part_score" validate:"min=0.00,max=16.00"`
	ReadingPartScore    float32 `json:"read_part_score" validate:"min=0.00,max=36.00"`
	StructPartScore     float32 `json:"struct_part_score" validate:"min=0.00,max=33.00"`
	VocabularyPartScore float32 `json:"vocabulary_part_score"  validate:"min=0.00,max=35.00"`
}
