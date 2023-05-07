package usecase

type (
	TodoOptions struct {
		limit  *int
		offset *int
		target *string
	}
	TodoOption func(*TodoOptions)
)

func applyTodoOptions(optionItems ...TodoOption) TodoOptions {
	result := TodoOptions{}
	for _, opt := range optionItems {
		opt(&result)
	}

	return result
}
func TodoLimit(limit int) TodoOption {
	return func(t *TodoOptions) {
		if limit != 0 {
			t.limit = &limit
		}
	}
}

func TodoOffset(offset int) TodoOption {
	return func(t *TodoOptions) {
		if offset != 0 {
			t.offset = &offset
		}
	}
}

func TodoTarget(target string) TodoOption {
	return func(t *TodoOptions) {
		if target != "" {
			t.target = &target
		}
	}
}
