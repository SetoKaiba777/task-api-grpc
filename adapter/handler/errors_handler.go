package handler

import (
	"go-challenger/core/domain"

	"google.golang.org/grpc/codes"
)

func GetStatus(e error) codes.Code {
	switch(e){
	case domain.ErrInvalidStatus:
		return codes.InvalidArgument
	case domain.ErrNotFoundTask:
		return codes.NotFound
	default:
		return codes.Internal
	}
}