package helper

import "github.com/rs/zerolog"

// LoggerWithJobContext adds the job name to the logger as context.
func LoggerWithJobContext(logger zerolog.Logger, job string) zerolog.Logger {
	return logger.With().Str("job", job).Logger()
}
