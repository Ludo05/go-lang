package SwitchStatement

const French = "french"
const Spanish = "spanish"

func SayHello(name, language string) string {
	if name == "" {
		name = "World"
		return LanguageSwitch(language) + name
	}
	return LanguageSwitch(language) + name
}

func LanguageSwitch(language string) (prefix string) {
	switch language {
	case French:
		prefix = "Bonjour "

	case Spanish:
		prefix = "Hola "

	default:
		prefix = "Hello "
	}
	return
}
