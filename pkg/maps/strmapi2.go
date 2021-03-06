package maps

type StringCopier interface {
    Copy() string
}

type StringComparer interface {
    // Compare should return a value less than 0 if the receiver is less than the given value,
    // 0 if equal, and a value greater than 0 if the receiver is greater than the given value.
    Compare(string) int
}

