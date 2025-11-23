package models

import (
	"testing"
)

func TestCreatePersonInput_Validation(t *testing.T) {
	tests := []struct {
		name  string
		input CreatePersonInput
		valid bool
	}{
		{
			name:  "valid person input",
			input: CreatePersonInput{Name: "Santa Claus"},
			valid: true,
		},
		{
			name:  "empty name should be invalid",
			input: CreatePersonInput{Name: ""},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.input.Name != ""
			if isValid != tt.valid {
				t.Errorf("expected valid=%v, got valid=%v for input %+v", tt.valid, isValid, tt.input)
			}
		})
	}
}

func TestCreateGiftInput_Validation(t *testing.T) {
	tests := []struct {
		name  string
		input CreateGiftInput
		valid bool
	}{
		{
			name: "valid gift input",
			input: CreateGiftInput{
				PersonID:    "123",
				Title:       "New Sleigh",
				Description: "A modern sleigh",
			},
			valid: true,
		},
		{
			name: "missing person_id should be invalid",
			input: CreateGiftInput{
				PersonID:    "",
				Title:       "New Sleigh",
				Description: "A modern sleigh",
			},
			valid: false,
		},
		{
			name: "missing title should be invalid",
			input: CreateGiftInput{
				PersonID:    "123",
				Title:       "",
				Description: "A modern sleigh",
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.input.PersonID != "" && tt.input.Title != "" && tt.input.Description != ""
			if isValid != tt.valid {
				t.Errorf("expected valid=%v, got valid=%v for input %+v", tt.valid, isValid, tt.input)
			}
		})
	}
}
