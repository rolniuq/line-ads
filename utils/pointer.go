package utils

func NewPointer[T any](value T) *T {
	return &value
}

func GetPointerValue[T any](pointer *T) T {
	var res T
	if pointer == nil {
		return res
	}

	return *pointer
}
