package input

type Handler struct {
}

func FindCommand(kp *KeyPressInfo) []string {
	if kp.GetName() == "left" {
		return []string{"move-point-left-char"}
	}
	if kp.GetName() == "right" {
		return []string{"move-point-right-char"}
	}

	return nil
}
