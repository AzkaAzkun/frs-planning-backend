package entity

type Course struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	ClassSettingID int64  `json:"class_setting_id"`
}

func (Course) TableName() string {
	return "courses"
}
