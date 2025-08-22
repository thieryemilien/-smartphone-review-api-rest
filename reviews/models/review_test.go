package models

import (
	"testing"
)

func NewReviewCMD(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

// Test_CreateReviewCMD_validate tests the validate method of CreateReviewCMD
func Test_withCorrectParams(t *testing.T) {
	r := NewReviewCMD(5, "Great product!")
	err := r.validate()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		t.Fail()
	}
}

func Test_shouldFailWithInvalidStars(t *testing.T) {
	r := NewReviewCMD(6, "Great product!")
	err := r.validate()
	if err == nil {
		t.Errorf("expected error for invalid stars, got nil")
		t.Fail()
	}
}
