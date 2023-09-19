package repository

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/pkg/database/postgres"
)

func Save(petition *models.Petition) error {
	err := postgres.Database.Create(petition).Error
	if err != nil {
		return err
	}
	return nil
}

// func Update(petition *models.Petition) error {
// 	err := postgres.Database.Model(&models.Petition{}).Where("id = ?", petition.ID).Update("status", petition.Status).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
