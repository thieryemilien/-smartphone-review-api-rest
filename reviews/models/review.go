package models

import (
	"errors"
	"strconv"
	"time"
)

const maxCommentLength = 500

// Review represents a structure for review
type Review struct {
	Id      int64
	Stars   int // from 1 to 5
	Comment string
	Date    time.Time
}

// CreateReviewCMD represents a command to create a review
type CreateReviewCMD struct {
	Stars   int    `json:"stars"` // from 1 to 5
	Comment string `json:"comment"`
}

// validate checks if the CreateReviewCMD has valid data
func (r *CreateReviewCMD) validate() error {

	// Validate stars
	if r.Stars < 1 || r.Stars > 5 {
		return errors.New("stars must be between 1 and 5")
	}

	// Validate comment
	if len(r.Comment) > maxCommentLength {
		return errors.New("comment must not exceed " + strconv.Itoa(maxCommentLength) + " characters")
	}

	return nil
}
