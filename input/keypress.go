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
	str := kpi.name
	if kpi.HyperMod {
		str = "Hyper-" + str
	}
	if kpi.SuperMod {
		str = "Super-" + str
	}
	if kpi.MetaMod {
		str = "Alt-" + str
	}
	if kpi.ShiftMod {
		str = "Shift-" + str
	}
	if kpi.CtrlMod {
		str = "Ctrl-" + str
	}
	return str
}
