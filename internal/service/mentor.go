package service

import (
	"fmt"
	"github.com/Tilvaldiyev/mentoring-app-backend/internal/model"
	"strings"
)

func (s *Service) GetMentorsList(searchMap map[string]string) ([]model.MentorResponse, error) {
	searchTerm, ok := searchMap["search_term"]
	if ok {
		term := strings.ReplaceAll(searchTerm, " ", "%")
		searchMap["search_term"] = term
	}

	mentors, err := s.repo.GetMentorsList(searchMap)
	if err != nil {
		return nil, fmt.Errorf("DB err: %w", err)
	}

	return mentors, nil
}
