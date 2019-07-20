package logging

import (
	"context"
	"os"

	"cloud.google.com/go/logging"
	"github.com/monzo/slog"
)

var severityMap = map[slog.Severity]logging.Severity{
	slog.CriticalSeverity: logging.Critical,
	slog.ErrorSeverity:    logging.Error,
	slog.WarnSeverity:     logging.Warning,
	slog.InfoSeverity:     logging.Notice,
	slog.DebugSeverity:    logging.Info,
	slog.TraceSeverity:    logging.Debug,
}

type stackdriverLogger struct {
	logger        *logging.Logger
	defaultLabels map[string]string
}

func (s *stackdriverLogger) Log(events ...slog.Event) {
	for _, event := range events {
		sev, ok := severityMap[event.Severity]
		if !ok {
			sev = logging.Default
		}

		s.logger.Log(logging.Entry{
			Severity: sev,
			Labels:   s.defaultLabels,
		})
	}
}

func (s *stackdriverLogger) Flush() error {
	return s.logger.Flush()
}

// Init sets up the default logger.
func Init(service string) {
	parent := "projects/lolibrary-180111"
	if envParent := os.Getenv("LOG_STACKDRIVER_PARENT"); envParent != "" {
		parent = envParent
	}

	logger, err := logging.NewClient(context.Background(), parent)
	if err != nil {
		panic(err)
	}

	slog.SetDefaultLogger(&stackdriverLogger{logger: logger.Logger(service)})
}
