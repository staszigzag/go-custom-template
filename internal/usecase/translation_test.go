package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/staszigzag/go-custom-template/internal/models"
	"github.com/staszigzag/go-custom-template/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func translation(t *testing.T) (*usecase.TranslationUseCase, *MocktranslationRepo, *MocktranslationWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMocktranslationRepo(mockCtl)
	webAPI := NewMocktranslationWebAPI(mockCtl)

	translation := usecase.NewTranslationUseCase(repo, webAPI)

	return translation, repo, webAPI
}

func TestHistory(t *testing.T) {
	t.Parallel()

	translation, repo, _ := translation(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, nil)
			},
			res: []models.Translation(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, errInternalServErr)
			},
			res: []models.Translation(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := translation.History(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestTranslate(t *testing.T) {
	t.Parallel()

	translation, repo, webAPI := translation(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				webAPI.EXPECT().Translate(models.Translation{}).Return(models.Translation{}, nil)
				repo.EXPECT().Store(context.Background(), models.Translation{}).Return(nil)
			},
			res: models.Translation{},
			err: nil,
		},
		{
			name: "web API error",
			mock: func() {
				webAPI.EXPECT().Translate(models.Translation{}).Return(models.Translation{}, errInternalServErr)
			},
			res: models.Translation{},
			err: errInternalServErr,
		},
		{
			name: "repo error",
			mock: func() {
				webAPI.EXPECT().Translate(models.Translation{}).Return(models.Translation{}, nil)
				repo.EXPECT().Store(context.Background(), models.Translation{}).Return(errInternalServErr)
			},
			res: models.Translation{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := translation.Translate(context.Background(), models.Translation{})

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
