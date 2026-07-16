package register

import (
	"errors"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/features/ecommerce"
)

type UseCase struct {
	userRepo     auth.UserRepository
	customerRepo ecommerce.CustomerRepository
	idProvider   common.IdProvider
}

func NewUseCase(userRepo auth.UserRepository, customerRepo ecommerce.CustomerRepository, provider common.IdProvider) *UseCase {
	return &UseCase{
		userRepo:     userRepo,
		customerRepo: customerRepo,
		idProvider:   provider,
	}
}

func (uc *UseCase) Handle(req *Request) (*Response, error) {

	if _, err := uc.userRepo.GetByEmail(req.Email); err == nil {

		return nil, errors.New("email already in use")
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	id, err := uc.idProvider.GetId()
	if err != nil {
		return nil, err
	}

	appModel := common.AppModel{PublicID: id}

	var user auth.User
	user, err = user.New(&appModel, hashedPassword, req.Email)
	if err != nil {
		return nil, err
	}

	// user fields are filled by orm
	if _, err := uc.userRepo.Create(&user); err != nil {
		// uc.db.Create(&user).Error{
		return nil, err
	}

	customer := ecommerce.Customer{
		AppModel:  appModel,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	}
	if _, err := uc.customerRepo.GetByEmail(req.Email); err == nil {

		return nil, errors.New("email already in use")
	}

	if _, err := uc.customerRepo.Create(&customer); err != nil {
		return nil, err
	}

	return &Response{
		PublicId:  customer.PublicID,
		Firstname: customer.FirstName,
		Lastname:  customer.LastName,
		Phone:     customer.Phone,
		Email:     customer.Email,
	}, nil
}
