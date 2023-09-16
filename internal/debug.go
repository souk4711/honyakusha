package internal

var (
	isDebugMode = false
)

func EnableDebugMode() {
	isDebugMode = true
}

func IsDebugMode() bool {
	return isDebugMode
}
