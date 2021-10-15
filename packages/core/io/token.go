package io

type Token int

const (
	// Percent is a delimiter and signals the start of a comment.
	//
	// Glyph: %
	//
	// Hexadecimal: 25
	Percent Token = iota

	// Solidus is a delimiter.
	//
	// Glyph: /
	//
	// Hexadecimal: 2F
	Solidus

	// LeftParenthesis is a delimiter.
	//
	// Glyph: (
	//
	// Hexadecimal: 28
	LeftParenthesis

	// RightParenthesis is a delimiter.
	//
	// Glyph: )
	//
	// Hexadecimal: 29
	RightParenthesis

	// LessThan is a delimiter.
	//
	// Glyph: <
	//
	// Hexadecimal: 3C
	LessThan

	// GreaterThan is a delimiter.
	//
	// Glyph: >
	//
	// Hexadecimal: 3E
	GreaterThan

	// LeftSquareBracket is a delimiter.
	//
	// Glyph: [
	//
	// Hexadecimal: 5B
	LeftSquareBracket

	// RightSquareBracket is a delimiter.
	//
	// Glyph: ]
	//
	// Hexadecimal: 5D
	RightSquareBracket

	// LeftCurlyBracket is a delimiter.
	//
	// Glyph: {
	//
	// Hexadecimal: 7B
	LeftCurlyBracket

	// RightCurlyBracket is a delimiter.
	//
	// Glyph: }
	//
	// Hexadecimal: 7D
	RightCurlyBracket

	// Null is white-space character.
	//
	// Hexadecimal: 0x00
	//
	// Rune: NUL
	Null

	// HorizontalTab is white-space character.
	//
	// Hexadecimal: 0x09
	//
	// Rune: HT
	HorizontalTab

	// LineFeed is white-space character.
	//
	// Hexadecimal: 0x0A
	//
	// Rune: LF
	LineFeed

	// FormFeed is white-space character.
	//
	// Hexadecimal: 0x0C
	//
	// Rune: FF
	FormFeed

	// CarriageReturn is white-space character.
	//
	// Hexadecimal: 0x0D
	//
	// Rune: CR
	CarriageReturn

	// Space is white-space character.
	//
	// Hexadecimal: 0x20
	//
	// Rune: SP
	Space

	// Whitespace separates syntactic constructs such as names and number from each other.
	// Any sequence of consecutive white-space characters are treated as one character.
	//
	// The following tokens are treated as white-space characters:
	// Null, HorizontalTab, LineFeed, FormFeed, CarriageReturn, Space
	Whitespace

	Unknown
)

const (
	NUL     rune = 0
	HT           = 9
	LF           = 10
	FF           = 12
	CR           = 13
	SP           = 32
	PERCENT      = '%'
)
