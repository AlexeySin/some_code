package dferr

import "errors"

var (
	ErrNoTrainLegs = errors.New("NO_TRAIN_LENGS_FOUND")
	ErrUserInput   = errors.New("INVALID_PRIORITY")
)
