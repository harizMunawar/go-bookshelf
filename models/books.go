package models

type Books struct {
	Id        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Year      int    `json:"year" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Summary   string `json:"summary" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	PageCount int    `json:"pageCount" binding:"required"`
	ReadPage  int    `json:"readPage"`
	Reading   bool   `json:"reading"`
	Finished  bool   `json:"finished"`
}
