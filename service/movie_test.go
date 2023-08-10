package service

import (
	"MovieApi/model"
	"MovieApi/repository/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_MovieService_GetMovieByID(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Movie
		expectedError  error
		onMovieRepo    func(mock *mocks.MockMovieRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onMovieRepo: func(mock *mocks.MockMovieRepo) {
			mock.EXPECT().GetMovieById(gomock.Any()).Return(model.Movie{
				ID:          1,
				Title:       "Pengabdi Setan 2 Comunion",
				Description: "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
				Rating:      7.0,
				Image:       "",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			}, nil).Times(1)
		},
		expectedResult: model.Movie{
			ID:          1,
			Title:       "Pengabdi Setan 2 Comunion",
			Description: "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
			Rating:      7.0,
			Image:       "",
			CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onMovieRepo: func(mock *mocks.MockMovieRepo) {
			mock.EXPECT().GetMovieById(gomock.Any()).Return(model.Movie{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onMovieRepo: func(mock *mocks.MockMovieRepo) {
			mock.EXPECT().GetMovieById(gomock.Any()).Return(model.Movie{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			movieRepo := mocks.NewMockMovieRepo(mockCtrl)

			if testCase.onMovieRepo != nil {
				testCase.onMovieRepo(movieRepo)
			}

			service := Service{
				repo: movieRepo,
			}

			res, err := service.GetMovieById(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}
