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
	ScorePercentage  float32 `json:"score_percentage" validated:"required , min=0 ,max=100"`
	CalPartScore     float32 `json:"cal_part_score" validated:"required"`
	ProblemPartScore float32 `json:"problem_part_score" validated:"required"`
	AppliedPartScore float32 `json:"applied_part_score" validated:"required"`
}

type GetSciAnalyticRequest struct {
	ScorePercentage  float32 `json:"score_percentage" validated:"required"`
	LessonPartScore  float32 `json:"lesson_part_score" validated:"required"`
	AppliedPartScore float32 `json:"applied_part_score" validated:"required"`
}

type GetEngAnalyticRequest struct {
	ScorePercentage     float32 `json:"score_percentage" validated:"required"`
	ExpressionPartScore float32 `json:"exp_part_score" validated:"required"`
	ReadingPartScore    float32 `json:"read_part_score" validated:"required"`
	StructPartScore     float32 `json:"struct_part_score" validated:"required"`
	VocabularyPartScore float32 `json:"vocabulary_part_score"  validated:"required"`
}
