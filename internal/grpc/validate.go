package grpc

import (
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (req *CreateTodoRequest) Validate() error {
	if req.GetContent() == "" {
		return status.Error(codes.InvalidArgument, "content is required")
	}

	return nil
}
