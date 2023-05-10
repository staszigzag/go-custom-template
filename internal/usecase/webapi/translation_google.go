package webapi

import (
	"github.com/staszigzag/go-custom-template/internal/models"
)

type TranslationWebAPI struct{}

func New() *TranslationWebAPI {
	return &TranslationWebAPI{}
}

func (t *TranslationWebAPI) Translate(translation models.Translation) (models.Translation, error) {
	return translation, nil
}
