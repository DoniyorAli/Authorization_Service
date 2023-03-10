package authorization

import (
	"UacademyGo/Blogpost/Authorization_Service/protogen/blogpost"
	"UacademyGo/Blogpost/Authorization_Service/security"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authService) Login(ctx context.Context, req *blogpost.LoginRequest) (*blogpost.TokenResponse, error) {
	log.Println("Login...")

	errAuth := errors.New("username or password is wrong")

	user, err := s.stg.GetUserByUsername(req.Username)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.Internal, errAuth.Error())
	}

	match, err := security.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "security.ComparePassword: %s", err.Error())
	}

	if !match {
		return nil, status.Errorf(codes.Internal, errAuth.Error())
	}

	m := map[string]interface{} {
		"user_id": user.Id,
		"username": user.Username,
	}
	
	tokenStr, err := security.GenerateJWT(m, time.Minute*10, s.cfg.SecretKey)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "security.GenerateJWT: %s", err.Error())
	}

	return &blogpost.TokenResponse{
		Token: tokenStr,
	}, nil
}

//*====================================================================================

func (s *authService) HasAcces(ctx context.Context, req *blogpost.TokenRequest) (*blogpost.HasAccesResponse, error) {
	log.Println("HasAccess...")

	result, err := security.ParseClaims(req.Token, s.cfg.SecretKey)
	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "security.ParseClaims: %s", err.Error()))
		return &blogpost.HasAccesResponse{
			User: nil,
			HasAcces: false,
		}, nil
	}
	
	log.Println(result.Username)

	user, err := s.stg.GetUserById(result.UserID)

	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "s.stg.GetUserById: %s", err.Error()))
		return &blogpost.HasAccesResponse{
			User: nil,
			HasAcces: false,
		}, nil
	}

	return &blogpost.HasAccesResponse{
		User: user,
		HasAcces: true,
	}, nil
}