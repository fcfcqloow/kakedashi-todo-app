package domain

type NullableString struct {
	Value  string
	IsNull bool
}

type NullableInteger struct {
	Value  int
	IsNull bool
}

func PtrToNullInt(p *int) NullableInteger {
	if p == nil {
		return NullableInteger{
			IsNull: true,
		}
	}

	return NullableInteger{
		IsNull: false,
		Value:  *p,
	}
}

func PtrToNullString(p *string) NullableString {
	if p == nil {
		return NullableString{
			IsNull: true,
		}
	}

	return NullableString{
		IsNull: false,
		Value:  *p,
	}
}

type Task struct {
	Id          string
	Value       string
	Topics      []string
	Priority    NullableInteger
	DoneDate    NullableInteger
	DeadLine    int
	CreatedDate int
}
