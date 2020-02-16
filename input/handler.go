package input

import "log"

var (
	commands map[string]*CommandBinding
)

type CommandBinding struct {
	Args []string
}

func makeBinding(args []string) *CommandBinding {
	return &CommandBinding{args}
}

func init() {
	commands = map[string]*CommandBinding{}
	commands["left"] = makeBinding([]string{"move-point-left-char"})
	commands["right"] = makeBinding([]string{"move-point-right-char"})
	commands["up"] = makeBinding([]string{"move-point-up-line"})
	commands["down"] = makeBinding([]string{"move-point-down-line"})
	commands["return"] = makeBinding([]string{"insert-text", "\n"})
	for _, key := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()=+-_/?\\|\"'<>,.~`[]{}" {
		commands[string(key)] = makeBinding([]string{"insert-text", string(key)})
	}
}

func FindCommand(kp *KeyPressInfo) []string {
	kpName := kp.GetName()
	if item, found := commands[kpName]; found {
		return item.Args
	}
	log.Printf("Key not found: %v\n", kpName)
	return nil
}
