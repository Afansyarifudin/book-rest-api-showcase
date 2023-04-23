package service

import (
	"book-rest-api-showcase/model"
	"book-rest-api-showcase/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookByID(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult model.Book
		wantError      bool
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{
				ID:       1,
				Namebook: "dev",
				Author:   "Afan",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			Namebook: "dev",
			Author:   "Afan",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookByID(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookByID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			ID:       1,
			Namebook: "dev",
			Author:   "Afan",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:       1,
				Namebook: "dev",
				Author:   "Afan",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			Namebook: "dev",
			Author:   "Afan",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			ID:       1,
			Namebook: "dev",
			Author:   "Afan",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			ID:       1,
			Namebook: "dev-new",
			Author:   "Afan-new",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{
				ID:       1,
				Namebook: "dev-new",
				Author:   "Afan-new",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:       1,
			Namebook: "dev-new",
			Author:   "Afan-new",
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			ID:       1,
			Namebook: "dev-new",
			Author:   "Afan-new",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "record not found",
		wantError: true,
		input: model.Book{
			ID:       1,
			Namebook: "dev-new",
			Author:   "Afan-new",
		},
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.UpdateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}

func Test_BookService_GetAllBook(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult []model.Book
		wantError      bool
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name: "success",
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBooks().Return([]model.Book{
				{
					ID:       1,
					Namebook: "book-afan-1",
					Author:   "afan1",
				},
				{
					ID:       2,
					Namebook: "book-afan-2",
					Author:   "afan2",
				},
			}, nil).Times(1)
		},
		expectedResult: []model.Book{
			{
				ID:       1,
				Namebook: "book-afan-1",
				Author:   "afan1",
			},
			{
				ID:       2,
				Namebook: "book-afan-2",
				Author:   "afan2",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBooks().Return([]model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetAllBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		expectedError error
		onBookRepo    func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(nil).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
