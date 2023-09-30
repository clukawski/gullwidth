package gullwidth

import (
	"testing"
	"unicode/utf8"
)

const fullwidthValidTestInput = "! \" # $ % & ' ( ) * + , - . / 0 1 2 3 4 5 6 7 8 9 : ; < = > ? @ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z [ \\ ] ^ _ ` a b c d e f g h i j k l m n o p q r s t u v w x y z { | } ~"
const fullwidthInvalidTestInput = fullwidthValidTestInput + " ł"

// TestFullWidth tests a valid input and ensures the correct number of characters are returned
func TestFullWidth(t *testing.T) {
	t.Run("Test valid fullwidth ASCII set with spaces", func(t *testing.T) {
		fullwidth, err := Fullwidth(fullwidthValidTestInput)
		if err != nil {
			t.Fatalf("unable to format valid test input: %s", err)
		}

		// Fullwidth characters are 2 runes each.
		fullwidthCount := utf8.RuneCountInString(fullwidth) / 2
		if fullwidthCount != utf8.RuneCountInString(fullwidthValidTestInput) {
			t.Fatalf(
				"output (%d) not the same size as input (%d)",
				fullwidthCount,
				utf8.RuneCountInString(fullwidthValidTestInput),
			)
		}
	})
}

// TestFullWidthInvalidInput tests we fail on invalid input
func TestFullWidthInvalidInput(t *testing.T) {
	t.Run("Test invvalid fullwidth ASCII set with spaces", func(t *testing.T) {
		_, err := Fullwidth(fullwidthInvalidTestInput)
		if err == nil {
			t.Fatalf("formatted invalid test input")
		}
	})
}
