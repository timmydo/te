package ui

import "github.com/timmydo/te/input"

import "fmt"

var (
	keyMap map[uint]*input.KeyPressInfo
)

func init() {
	keyMap = make(map[uint]*input.KeyPressInfo)
	keyMap[0x61] = input.NewKeyPressInfo("a")
	keyMap[0xff51] = input.NewKeyPressInfo("left")
	keyMap[0xff52] = input.NewKeyPressInfo("up")
	keyMap[0xff53] = input.NewKeyPressInfo("right")
	keyMap[0xff54] = input.NewKeyPressInfo("down")
	keyMap[0xff55] = input.NewKeyPressInfo("pageup")
	keyMap[0xff56] = input.NewKeyPressInfo("pagedown")
	keyMap[0xff0d] = input.NewKeyPressInfo("return")

	keyMap[0x0020] = input.NewKeyPressInfo("space")

	keyMap[0xff50] = input.NewKeyPressInfo("home")
	keyMap[0xff57] = input.NewKeyPressInfo("end")

	keyMap[0xff08] = input.NewKeyPressInfo("backspace")
	keyMap[0xff09] = input.NewKeyPressInfo("tab")
	keyMap[0xff1b] = input.NewKeyPressInfo("escape")
	keyMap[0xffff] = input.NewKeyPressInfo("delete")

	for i, _ := range "123456789012" {
		keyMap[0xffbe+uint(i)] = input.NewKeyPressInfo(fmt.Sprintf("f%d", i+1))
	}
	for i, c := range "[\\]^_`\"" {
		keyMap[0x005b+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range "{|}~" {
		keyMap[0x007b+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range ":;<=>?@" {
		keyMap[0x003a+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range "!\"#$%&'()*+,-./" {
		keyMap[0x0021+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range "0123456789" {
		keyMap[0x0030+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range "abcdefghijklmnopqrstuvwxyz" {
		keyMap[0x0061+uint(i)] = input.NewKeyPressInfo(string(c))
	}
	for i, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		keyMap[0x0041+uint(i)] = input.NewKeyPressInfo(string(c))
	}
}
