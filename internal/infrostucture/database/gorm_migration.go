package database

import (
	"DataHub/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = db.Exec(`
		DO $$ 
		DECLARE
			r RECORD;
		BEGIN
			FOR r IN (
				SELECT DISTINCT table_name 
				FROM information_schema.columns 
				WHERE table_schema = 'public' 
				  AND data_type = 'uuid'
			) LOOP
				EXECUTE 'DROP TABLE IF EXISTS "' || r.table_name || '" CASCADE';
			END LOOP;
		END $$;
	`).Error

	err = db.AutoMigrate(
		&domain.Branch{},
		&domain.Department{},
		&domain.Position{},
		&domain.Folder{},
		&domain.Test{},
		&domain.Question{},
		&domain.Chat{},
		&domain.FormSession{},
		&domain.SchemaMigration{},
		&domain.GradingLog{},
		&domain.HR{},
		&domain.Mentor{},
		&domain.WelcomeContent{},
		&domain.TrainingTemplate{},
		&domain.SubDepartment{},
		&domain.Employee{},
		&domain.Option{},
		&domain.TestQuestion{},
		&domain.TestSession{},
		&domain.User{},
		&domain.ChatMember{},
		&domain.File{},
		&domain.Message{},
		&domain.Reaction{},
		&domain.Attempt{},
		&domain.Onboarding{},
		&domain.WelcomeProgress{},
		&domain.TrainingProgress{},
		&domain.MentorAssessment{},
		&domain.Task{},
		&domain.Article{},
		&domain.ArticleMedia{},
		&domain.Comment{},
		&domain.Announcement{},
		&domain.TestAssignment{},
		&domain.Answer{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
