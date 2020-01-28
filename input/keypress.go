package input

type KeyPressInfo struct {
	CtrlMod  bool
	ShiftMod bool
	MetaMod  bool
	SuperMod bool
	HyperMod bool

	character string
}

func NewKeyPressInfo(ch string) *KeyPressInfo {
	kp := new(KeyPressInfo)
	kp.character = ch
	return kp
}
