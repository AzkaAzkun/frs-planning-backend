package repository

import (
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	ClassSettingRepository interface {
		Create(ctx context.Context, tx *gorm.DB, classsetting entity.ClassSettings) (entity.ClassSettings, error)
		Clone(ctx context.Context, tx *gorm.DB, userid uuid.UUID, classsettingid uuid.UUID) (entity.ClassSettings, error)
	}

	classSettingRepository struct {
		db *gorm.DB
	}
)

func NewClassSettingRepository(db *gorm.DB) ClassSettingRepository {
	return &classSettingRepository{db}
}

func (r *classSettingRepository) Create(ctx context.Context, tx *gorm.DB, classsetting entity.ClassSettings) (entity.ClassSettings, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&classsetting).Error; err != nil {
		return entity.ClassSettings{}, err
	}
	return classsetting, nil
}

func (r *classSettingRepository) Clone(ctx context.Context, tx *gorm.DB, userID uuid.UUID, classSettingID uuid.UUID) (entity.ClassSettings, error) {
	if tx == nil {
		tx = r.db
	}

	var originalClassSetting entity.ClassSettings
	if err := tx.WithContext(ctx).First(&originalClassSetting, "id = ?", classSettingID).Error; err != nil {
		return entity.ClassSettings{}, err
	}

	newClassSetting := entity.ClassSettings{

		UserID:     userID,
		Permission: "PRIVATE",
		// Name and Status fields removed as they do not exist in the database schema

		Name:       originalClassSetting.Name,
		UserID:     userID,
		Permission: "PRIVATE",
		Status:     "CLONE",

	}

	if err := tx.WithContext(ctx).Create(&newClassSetting).Error; err != nil {
		return entity.ClassSettings{}, err
	}

	var courses []entity.Course
	if err := tx.WithContext(ctx).Where("class_setting_id = ?", classSettingID).Find(&courses).Error; err != nil {
		return entity.ClassSettings{}, err
	}

	courseIDMap := make(map[uuid.UUID]uuid.UUID)

	for _, course := range courses {
		newCourse := entity.Course{
			Name:           course.Name,
			ClassSettingID: newClassSetting.ID,
		}

		if err := tx.WithContext(ctx).Create(&newCourse).Error; err != nil {
			return entity.ClassSettings{}, err
		}

		courseIDMap[course.ID] = newCourse.ID

		var classes []entity.Class
		if err := tx.WithContext(ctx).Where("course_id = ?", course.ID).Find(&classes).Error; err != nil {
			return entity.ClassSettings{}, err
		}

		for _, class := range classes {
			newClass := entity.Class{
				Lecturer:      class.Lecturer,
				CourseID:      newCourse.ID,
				ClassSchedule: class.ClassSchedule,
				Priority:      class.Priority,
			}

			if err := tx.WithContext(ctx).Create(&newClass).Error; err != nil {
				return entity.ClassSettings{}, err
			}
		}
	}

	return newClassSetting, nil
}
