package models

//StudentBook - student dao
type StudentBook struct {
	Student *Student `json:"student,omitempty"`
	Books   *[]Book  `json:"books,omitempty"`
}
