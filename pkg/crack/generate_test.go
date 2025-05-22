package crack

import (
	"fmt"
	"github.com/lightningnetwork/lnd/aezeed"
	"golang.org/x/exp/slices"
	"testing"
)

var fullMnemonic = []string{
	"above", "bronze", "era", "decade", "crane", "fossil", "hand", "tomato",
	"midnight", "entry", "bridge", "know", "fat", "power", "vintage", "uncle",
	"dinner", "carry", "train", "hollow", "parrot", "burger", "anchor", "clown",
}

func TestFindSingleMissingWord(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test")
	}
	for i := 0; i < len(fullMnemonic); i++ {
		t.Run(fmt.Sprintf("pos_%d", i), func(t *testing.T) {
			partial := append([]string{}, fullMnemonic[:i]...)
			partial = append(partial, fullMnemonic[i+1:]...)
			settings := MnemonicSettings{
				Wordlist:       aezeed.DefaultWordList,
				Password:       "",
				Mnemonic:       partial,
				MnemonicLength: len(fullMnemonic),
			}
			done := make(chan struct{})
			var result Result
			FindMissingWords(settings, func(r Result) {
				result = r
				if r.CipherSeed != nil || r.Exhausted() {
					close(done)
				}
			})
			<-done
			if result.CipherSeed == nil {
				t.Fatalf("position %d: missing word not found", i)
			}
			mnemonic, err := result.CipherSeed.ToMnemonic([]byte(""))
			if err != nil {
				t.Fatalf("position %d: %v", i, err)
			}
			if !slices.Equal(mnemonic[:], fullMnemonic) {
				t.Fatalf("position %d: recovered mnemonic mismatch", i)
			}
			if len(result.Highlight) != 1 || result.Highlight[0] != i {
				t.Fatalf("position %d: highlight %v", i, result.Highlight)
			}
		})
	}
}

func TestFindTwoMissingWords(t *testing.T) {
	if testing.Short() {
		t.Skip("slow test")
	}
	for i := 0; i < len(fullMnemonic); i++ {
		for j := i + 1; j < len(fullMnemonic); j++ {
			t.Run(fmt.Sprintf("pos_%d_%d", i, j), func(t *testing.T) {
				partial := append([]string{}, fullMnemonic[:i]...)
				partial = append(partial, fullMnemonic[i+1:j]...)
				partial = append(partial, fullMnemonic[j+1:]...)
				settings := MnemonicSettings{
					Wordlist:       aezeed.DefaultWordList,
					Password:       "",
					Mnemonic:       partial,
					MnemonicLength: len(fullMnemonic),
				}
				done := make(chan struct{})
				var result Result
				FindMissingWords(settings, func(r Result) {
					result = r
					if r.CipherSeed != nil || r.Exhausted() {
						close(done)
					}
				})
				<-done
				if result.CipherSeed == nil {
					t.Fatalf("positions %d,%d: missing words not found", i, j)
				}
				mnemonic, err := result.CipherSeed.ToMnemonic([]byte(""))
				if err != nil {
					t.Fatalf("positions %d,%d: %v", i, j, err)
				}
				if !slices.Equal(mnemonic[:], fullMnemonic) {
					t.Fatalf("positions %d,%d: recovered mnemonic mismatch", i, j)
				}
				if len(result.Highlight) != 2 || result.Highlight[0] != i || result.Highlight[1] != j {
					t.Fatalf("positions %d,%d: highlight %v", i, j, result.Highlight)
				}
			})
		}
	}
}
