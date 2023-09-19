package service

import (
	"github.com/catness812/e-petitions-project/internal/models"
	"github.com/catness812/e-petitions-project/internal/repository"
)

func CreateNew(petition models.Petition) (uint, error) {
	if err := repository.Save(&petition); err != nil {
		return 0, err
	} else {
		return petition.ID, nil
	}
}
