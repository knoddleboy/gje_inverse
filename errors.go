package gjeinverse

type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	ErrNegativeThreads   = Error{"mat: negative number of threads"}
	ErrNegativeDimension = Error{"mat: negative dimension"}
	ErrSingularMatrix    = Error{"mat: inverse does not exist"}
	ErrFailedToCompute   = Error{"mat: failed to compute the inverse matrix"}
)
