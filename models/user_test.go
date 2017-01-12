// models.user_test.go

package models

import (
	"testing"

)

// Test the validity of different combinations of username/password
func TestUserValidity(t *testing.T) {
	if !IsUserValid("user1", "pass1") {
		t.Fail()
	}

	if IsUserValid("user2", "pass1") {
		t.Fail()
	}

	if IsUserValid("user1", "") {
		t.Fail()
	}

	if IsUserValid("", "pass1") {
		t.Fail()
	}

	if IsUserValid("User1", "pass1") {
		t.Fail()
	}
}

// Test if a new user can be registered with valid username/password
func TestValidUserRegistration(t *testing.T) {
	SaveLists()

	u, err := RegisterNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	RestoreLists()
}

// Test that a new user cannot be registered with invalid username/password
func TestInvalidUserRegistration(t *testing.T) {
	SaveLists()

	// Try to register a user with a used username
	u, err := RegisterNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	// Try to register with a blank password
	u, err = RegisterNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	RestoreLists()
}

// Test the function that checks for username availability
func TestUsernameAvailability(t *testing.T) {
	SaveLists()

	// This username should be available
	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	// This username should not be available
	if isUsernameAvailable("user1") {
		t.Fail()
	}

	// Register a new user
	RegisterNewUser("newuser", "newpass")

	// This newly registered username should not be available
	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	RestoreLists()
}
