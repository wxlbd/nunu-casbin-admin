package types

type MenuType int8

const (
	MenuTypeMenu MenuType = iota + 1
	MenuTypeLink
	MenuTypeIframe
	MenuTypeButton
)
