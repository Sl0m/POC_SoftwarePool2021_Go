// Code generated by entc, DO NOT EDIT.

package ent

import (
	"go-message/ent/message"
	"go-message/ent/room"
	"go-message/ent/schema"
	"go-message/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescMessage is the schema descriptor for message field.
	messageDescMessage := messageFields[1].Descriptor()
	// message.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	message.MessageValidator = messageDescMessage.Validators[0].(func(string) error)
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[2].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}