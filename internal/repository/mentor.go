package repository

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
)

func (r *Repository) GetMentorsList(searchMap map[string]string) ([]model.MentorResponse, error) {
	var mentors []model.MentorResponse

	var queryParam string
	var queryVals []interface{}
	for key, value := range searchMap {
		if key == "search_term" {
			queryParam += "first_name || ' ' || last_name || title LIKE ? and "
			queryVals = append(queryVals, fmt.Sprintf("%%%s%%", value))
		} else {
			queryParam += key + " = ? and "
			queryVals = append(queryVals, value)
		}
	}

	queryParam += "u.user_type_id = 3"

	err := r.DB.Select(`u.id, u.last_name, u.first_name, u.email, u.title, u.info,
		l.name as language, c.name as country, lev.name as level, e.name as expertise`).
		Table(model.Users{}.TableName() + " u").
		Joins("inner join expertise_user eu on eu.id = u.id").
		Joins("inner join expertise e on e.id = eu.expertise_id").
		Joins("inner join language l on l.id = u.language_id").
		Joins("inner join country c on c.id = u.country_id").
		Joins("inner join level lev on lev.id = u.level_id").
		Where(queryParam, queryVals...).Scan(&mentors).Error
	if err != nil {
		return mentors, fmt.Errorf("getting mentors list err: %w", err)
	}


	return mentors, nil
}
