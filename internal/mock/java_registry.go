package mock

type JavaRegistry struct {
	Registry map[string]string
}

func NewJavaRegistry() *JavaRegistry {
	newRegistry := &JavaRegistry{
		Registry: make(map[string]string),
	}

	newRegistry.Registry["jdk17"] = "C:\\Users\\steve\\.sdk\\java\\jdk17"
	newRegistry.Registry["jdk21"] = "C:\\Users\\steve\\.sdk\\java\\jdk21"

	return newRegistry
}
