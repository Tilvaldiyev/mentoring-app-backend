package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/config"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// NewDB - connecting to the DB, migrating
func NewDB(cfgDB *config.DB) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s TimeZone=%s",
		cfgDB.Host, strconv.Itoa(cfgDB.Port), cfgDB.Username, cfgDB.Name, cfgDB.SSLMode, os.Getenv("POSTGRES_PASSWORD"), cfgDB.TimeZone)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("openning conn err: %w", err)
	}

	//pgDB, err := db.DB()
	//if err != nil {
	//	return nil, fmt.Errorf("connecting to pg db err: %w", err)
	//}

	//pgDB.SetMaxIdleConns(cfgDB.MaxIdleConns)
	//pgDB.SetMaxOpenConns(cfgDB.MaxOpenConns)
	//pgDB.SetConnMaxLifetime(cfgDB.ConnMaxLifetime * time.Second)
	//pgDB.SetConnMaxIdleTime(cfgDB.ConnMaxIdleTime * time.Second)
	manager := Repository{DB: db}
	if cfgDB.Migrate {
		err = manager.autoMigrate()
		if err != nil {
			return nil, fmt.Errorf("migration err: %w", err)
		}
	}

	err = manager.initIndexes()
	if err != nil {
		return nil, fmt.Errorf("initiating indexes err: %w", err)
	}

	err = manager.setBaseValues()
	if err != nil {
		return nil, fmt.Errorf("set base values err: %w", err)
	}

	return db, nil
}

// autoMigrate - migration
func (r *Repository) autoMigrate() error {
	return r.DB.AutoMigrate(
		&model.Language{},
		&model.UserType{},
		&model.Country{},
		&model.Expertise{},
		&model.Users{},
		&model.ExpertiseUser{},
		&model.Posts{},
		&model.Article{},
		&model.Level{},
		&model.ExpertiseArticle{},
	)

}

// initIndexes - creating indexes
func (r *Repository) initIndexes() error {
	languageTable := model.Language{}.TableName()
	sql := fmt.Sprintf("create unique index if not exists language_id_uindex on %s (id)", languageTable)
	err := r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", languageTable, err)
	}

	userTypeTable := model.UserType{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists userType_id_uindex on %s (id)", userTypeTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", userTypeTable, err)
	}

	countryTable := model.Country{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists country_id_uindex on %s (id)", countryTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", countryTable, err)
	}

	expertiseTable := model.Expertise{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists expertise_id_uindex on %s (id)", expertiseTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", expertiseTable, err)
	}

	usersTable := model.Users{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists users_id_uindex on %s (id)", usersTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", usersTable, err)
	}

	expertiseUserTable := model.ExpertiseUser{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists expertise_user_id_uindex on %s (id)", expertiseUserTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", expertiseUserTable, err)
	}

	postsTable := model.Posts{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists posts_id_uindex on %s (id)", postsTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", postsTable, err)
	}

	articleTable := model.Article{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists article_id_uindex on %s (id)", articleTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", articleTable, err)
	}

	levelTable := model.Level{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists level_id_uindex on %s (id)", levelTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", levelTable, err)
	}

	expertiseArticleTable := model.ExpertiseArticle{}.TableName()
	sql = fmt.Sprintf("create unique index if not exists expertise_article_id_uindex on %s (id)", expertiseArticleTable)
	err = r.DB.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("creating index for table %s err: %w", expertiseArticleTable, err)
	}

	return nil
}

func (r *Repository) setBaseValues() error {
	var count int64
	if err := r.DB.Table(model.Country{}.TableName()).Count(&count).Error; err != nil {
		return fmt.Errorf("get num of rows err: %w", err)
	}
	if count == 0 {
		countries := make([]model.Country, 0, 28)
		countries = append(countries, model.Country{Name: "Afghanistan"})
		countries = append(countries, model.Country{Name: "Albania"})
		countries = append(countries, model.Country{Name: "Argentina"})
		countries = append(countries, model.Country{Name: "Armenia"})
		countries = append(countries, model.Country{Name: "Australia"})
		countries = append(countries, model.Country{Name: "Austria"})
		countries = append(countries, model.Country{Name: "Azerbaijan"})
		countries = append(countries, model.Country{Name: "Belarus"})
		countries = append(countries, model.Country{Name: "Belgium"})
		countries = append(countries, model.Country{Name: "Brazil"})
		countries = append(countries, model.Country{Name: "Bulgaria"})
		countries = append(countries, model.Country{Name: "Canada"})
		countries = append(countries, model.Country{Name: "China"})
		countries = append(countries, model.Country{Name: "Czech Republic"})
		countries = append(countries, model.Country{Name: "Denmark"})
		countries = append(countries, model.Country{Name: "Egypt"})
		countries = append(countries, model.Country{Name: "France"})
		countries = append(countries, model.Country{Name: "Georgia"})
		countries = append(countries, model.Country{Name: "Germany"})
		countries = append(countries, model.Country{Name: "Greece"})
		countries = append(countries, model.Country{Name: "Japan"})
		countries = append(countries, model.Country{Name: "Kazakhstan"})
		countries = append(countries, model.Country{Name: "Kyrgyzstan"})
		countries = append(countries, model.Country{Name: "Mexico"})
		countries = append(countries, model.Country{Name: "Portugal"})
		countries = append(countries, model.Country{Name: "Spain"})
		countries = append(countries, model.Country{Name: "Sweden"})
		countries = append(countries, model.Country{Name: "Uzbekistan"})

		err := r.DB.Create(&countries).Error
		if err != nil {
			return fmt.Errorf("inserting err: %w", err)
		}
	}

	if err := r.DB.Table(model.Expertise{}.TableName()).Count(&count).Error; err != nil {
		return fmt.Errorf("get num of rows err: %w", err)
	}
	if count == 0 {
		expertises := make([]model.Expertise, 0, 4)
		expertises = append(expertises, model.Expertise{Name: "Design"})
		expertises = append(expertises, model.Expertise{Name: "Marketing"})
		expertises = append(expertises, model.Expertise{Name: "Product Management"})
		expertises = append(expertises, model.Expertise{Name: "Software Development"})

		err := r.DB.Create(&expertises).Error
		if err != nil {
			return fmt.Errorf("inserting err: %w", err)
		}
	}

	if err := r.DB.Table(model.Language{}.TableName()).Count(&count).Error; err != nil {
		return fmt.Errorf("get num of rows err: %w", err)
	}
	if count == 0 {
		languages := make([]model.Language, 0, 13)
		languages = append(languages, model.Language{Name: "English"})
		languages = append(languages, model.Language{Name: "French"})
		languages = append(languages, model.Language{Name: "German"})
		languages = append(languages, model.Language{Name: "Italian"})
		languages = append(languages, model.Language{Name: "Japanese"})
		languages = append(languages, model.Language{Name: "Korean"})
		languages = append(languages, model.Language{Name: "Portuguese"})
		languages = append(languages, model.Language{Name: "Russian"})
		languages = append(languages, model.Language{Name: "Spanish"})
		languages = append(languages, model.Language{Name: "Turkish"})
		languages = append(languages, model.Language{Name: "Kazakh"})
		languages = append(languages, model.Language{Name: "Ukrainian"})
		languages = append(languages, model.Language{Name: "Chinese"})

		err := r.DB.Create(&languages).Error
		if err != nil {
			return fmt.Errorf("inserting err: %w", err)
		}
	}

	if err := r.DB.Table(model.Level{}.TableName()).Count(&count).Error; err != nil {
		return fmt.Errorf("get num of rows err: %w", err)
	}
	if count == 0 {
		levels := make([]model.Level, 0, 4)
		levels = append(levels, model.Level{Name: "Entry Level"})
		levels = append(levels, model.Level{Name: "Junior"})
		levels = append(levels, model.Level{Name: "Middle"})
		levels = append(levels, model.Level{Name: "Senior"})

		err := r.DB.Create(&levels).Error
		if err != nil {
			return fmt.Errorf("inserting err: %w", err)
		}
	}

	if err := r.DB.Table(model.UserType{}.TableName()).Count(&count).Error; err != nil {
		return fmt.Errorf("get num of rows err: %w", err)
	}
	if count == 0 {
		types := make([]model.UserType, 0, 2)
		types = append(types, model.UserType{Name: "Mentor"})
		types = append(types, model.UserType{Name: "Mentee"})

		err := r.DB.Create(&types).Error
		if err != nil {
			return fmt.Errorf("inserting err: %w", err)
		}
	}

	return nil
}
