package usecase

import (
	"context"
	"fmt"

	"github.com/staszigzag/go-custom-template/internal/models"
)

//go:generate mockgen -source=./translation.go -destination=./mocks_test.go -package=usecase_test
type translationRepo interface {
	Store(context.Context, models.Translation) error
	GetHistory(context.Context) ([]models.Translation, error)
}

type translationWebAPI interface {
	Translate(models.Translation) (models.Translation, error)
}

type TranslationUseCase struct {
	repo   translationRepo
	webAPI translationWebAPI
}

func NewTranslationUseCase(r translationRepo, w translationWebAPI) *TranslationUseCase {
	return &TranslationUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *TranslationUseCase) History(ctx context.Context) ([]models.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}

func (uc *TranslationUseCase) Translate(ctx context.Context, t models.Translation) (models.Translation, error) {
	translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return models.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	err = uc.repo.Store(context.Background(), translation)
	if err != nil {
		return models.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return translation, nil
}
