package handlers

import (
	"context"
	usersrc "goapi/internal/userSrc"
	"goapi/internal/web/users"
)

type UserHandler struct {
	service usersrc.UserService
}

func NewUserHandler(s usersrc.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allusers, err := u.service.GetUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, usr := range allusers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

// PatchUsersUserId implements users.StrictServerInterface.
func (u *UserHandler) PatchUsersUserId(ctx context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	taskId := request.UserId
	userToUpdate, err := u.service.GetUserByID(taskId)
	if err != nil {
		return nil, err
	}
	userToUpdate.Email = *request.Body.Email
	userToUpdate.Password = *request.Body.Password
	response := users.PatchUsersUserId200JSONResponse{
		Id:       &userToUpdate.ID,
		Email:    &userToUpdate.Email,
		Password: &userToUpdate.Password,
	}
	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	createdUser, err := u.service.CreateUser(*userRequest.Email, *userRequest.Password)
	if err != nil {
		return nil, err
	}
	return users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}, nil
}

// DeleteUsersUserId implements users.StrictServerInterface.
func (u *UserHandler) DeleteUsersUserId(ctx context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	userId := request.UserId
	err := u.service.DeleteUser(string(userId))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersUserId204Response{}, nil
}
