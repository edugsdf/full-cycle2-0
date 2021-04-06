package application_test

import (
	"testing"

	mock_application "github.com/edugsdf/full-cycle2-0/tree/main/docker/golang/hexagonal-aluno/mocks"
	"github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTime()

}
