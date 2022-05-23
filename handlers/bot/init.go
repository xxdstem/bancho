package bot

func HandleMessage(message string) string {
	if message == "!test" {
		return "bonk"
	} else {
		return "niggas!"
	}
}
