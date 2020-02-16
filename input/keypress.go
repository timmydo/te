package input

type KeyPressInfo struct {
	CtrlMod  bool
	ShiftMod bool
	MetaMod  bool
	SuperMod bool
	HyperMod bool

	name string
}

func NewKeyPressInfo(ch string) *KeyPressInfo {
	kp := new(KeyPressInfo)
	kp.name = ch
	return kp
}

func (kpi *KeyPressInfo) GetName() string {
	return kpi.name
}
