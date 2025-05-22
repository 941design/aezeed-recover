package ui

import (
	"testing"

	"github.com/lightningnetwork/lnd/aezeed"
	"github.com/rivo/tview"
	"golang.org/x/exp/slices"
)

func newTestSeedView() *SeedView {
	theme := UniversalTheme()
	config := &Config{
		Wordlist:        aezeed.DefaultWordList,
		DefaultPassword: "default",
		MnemonicLength:  24,
	}
	sv := NewSeedView(theme, config, func(p tview.Primitive) *tview.Application { return nil })
	sv.SetOnMnemonicChangedFunc(func([]string, bool) {})
	return sv
}

func TestSeedViewSetAndGetWords(t *testing.T) {
	sv := newTestSeedView()
	words := []string{
		"abandon", "ability", "able", "about", "above", "absent", "absorb", "abstract", "absurd", "abuse", "access", "accident",
		"account", "accuse", "achieve", "acid", "acoustic", "acquire", "across", "act", "action", "actor", "actress", "actual",
	}
	sv.SetWords(words)
	got := sv.GetWords()
	if !slices.Equal(got, words) {
		t.Fatalf("expected %v, got %v", words, got)
	}
}

func TestSeedViewGetPassword(t *testing.T) {
	sv := newTestSeedView()
	sv.passwordInput.SetText("")
	if pw := sv.GetPassword(); pw != sv.config.DefaultPassword {
		t.Fatalf("expected default password %q, got %q", sv.config.DefaultPassword, pw)
	}
	sv.passwordInput.SetText("secret")
	if pw := sv.GetPassword(); pw != "secret" {
		t.Fatalf("expected password 'secret', got %q", pw)
	}
}

func TestSeedViewGetMnemonic(t *testing.T) {
	sv := newTestSeedView()
	words := []string{
		"abandon", "ability", "able", "about", "above", "absent", "absorb", "abstract", "absurd", "abuse", "access", "accident",
		"account", "accuse", "achieve", "acid", "acoustic", "acquire", "across", "act", "action", "actor", "actress", "actual",
	}
	sv.SetWords(words)
	mnemonic, err := sv.GetMnemonic()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !slices.Equal(mnemonic[:], words) {
		t.Fatalf("expected mnemonic %v, got %v", words, mnemonic[:])
	}
	sv = newTestSeedView()
	sv.SetWords(words[:23])
	if _, err := sv.GetMnemonic(); err == nil {
		t.Fatalf("expected error for incomplete mnemonic")
	}
}

func TestWordCompleter(t *testing.T) {
	wc := WordCompleter([]string{"foo", "bar", "baz"})
	got := wc("ba")
	expected := []string{"bar", "baz"}
	if !slices.Equal(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
