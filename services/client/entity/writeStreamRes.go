package entity

type CreateStreamRes struct {
	progress uint8
	err      error
}

// Setter method for progress field
func (w *CreateStreamRes) SetProgress(progress uint8) {
	w.progress = progress
}

// Getter method for progress field
func (w *CreateStreamRes) GetProgress() uint8 {
	return w.progress
}

func (w *CreateStreamRes) IsError() bool {
	return w.err != nil
}

// Setter method for error field
func (w *CreateStreamRes) SetError(err error) {
	w.err = err
}

// Getter method for error field
func (w *CreateStreamRes) GetError() error {
	return w.err
}
