package utils

type Questionnaire struct {
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Title string `json:"title"`
	Type  string `json:"type"`
}
