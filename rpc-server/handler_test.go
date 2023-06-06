package main

import (
	"testing"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/stretchr/testify/assert"
)

func TestValidateSendRequest(t *testing.T) {
	validRequest := &rpc.SendRequest{
		Message: &rpc.Message{
			Chat:   "user1:user2",
			Sender: "user1",
			Text:   "",
		},
	}

	invalidRequest := &rpc.SendRequest{
		Message: &rpc.Message{
			Chat:   "invalid",
			Sender: "user1",
			Text:   "",
		},
	}

	err := validateSendRequest(validRequest)
	assert.NoError(t, err, "Expected no error for a valid request")

	err = validateSendRequest(invalidRequest)
	assert.Error(t, err, "Expected an error for an invalid request")
	assert.EqualError(t, err, "invalid Chat ID 'invalid', should be in the format of user1:user2")
}

func TestGetRoomID(t *testing.T) {
	roomID, err := getRoomID("user1:user2")
	assert.NoError(t, err, "Expected no error for a valid chat")

	expectedRoomID := "user1:user2"
	assert.Equal(t, expectedRoomID, roomID, "Expected the room ID to be user1:user2")

	roomID, err = getRoomID("invalid")
	assert.Error(t, err, "Expected an error for an invalid chat")
	assert.EqualError(t, err, "invalid Chat ID 'invalid', should be in the format of user1:user2")
	assert.Empty(t, roomID, "Expected an empty room ID for an invalid chat")
}