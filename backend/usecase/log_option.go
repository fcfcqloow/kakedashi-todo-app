package usecase

type (
	LogOptions struct {
		year  int
		month int
	}
	LogOption func(*LogOptions)
)

func LogYear(year int) LogOption {
	return func(l *LogOptions) {
		l.year = year
	}
}

func LogMonth(month int) LogOption {
	return func(l *LogOptions) {
		l.month = month
	}
}
