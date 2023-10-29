package user

import (
	"context"
	pb "gRPC_Assignment/proto"
	"testing"
)

func TestUserService_GetUserById(t *testing.T) {
	// Create a test instance of the UserService
	userService := NewUserService()

	// Define test cases
	testCases := []struct {
		name       string
		request    *pb.UserRequest
		expected   *pb.User
		expectErr  bool
		errMessage string
	}{
		{
			name: "Valid User ID",
			request: &pb.UserRequest{
				Id: 1,
			},
			expected: &pb.User{
				Id:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   1234567890,
				Height:  5.8,
				Married: true,
			},
			expectErr:  false,
			errMessage: "",
		},
		{
			name: "Invalid User ID",
			request: &pb.UserRequest{
				Id: 2,
			},
			expected:   nil,
			expectErr:  true,
			errMessage: "user not found",
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := userService.GetUserById(context.Background(), tc.request)

			// Check if the error condition matches the expectation
			if (err != nil) != tc.expectErr || (err != nil && err.Error() != tc.errMessage) {
				t.Errorf("Expected error: %v, got: %v", tc.errMessage, err)
			}

			// Check if the returned user matches the expectation
			if user == nil && tc.expected != nil {
				t.Errorf("Expected user to be %v, got nil", tc.expected)
			} else if user != nil && (user.Id != tc.expected.Id || user.Fname != tc.expected.Fname || user.City != tc.expected.City || user.Phone != tc.expected.Phone || user.Height != tc.expected.Height || user.Married != tc.expected.Married) {
				t.Errorf("Expected user: %v, got: %v", tc.expected, user)
			}
		})
	}
}

// Manual implementation of UserService_GetUsersByIdsServer for testing
type testServer struct {
	data []*pb.User
}

func (t *testServer) Context() context.Context {
	// This method can be implemented according to the actual requirements
	return context.Background()
}

func (t *testServer) Send(user *pb.User) error {
	t.data = append(t.data, user)
	return nil
}
