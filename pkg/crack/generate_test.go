package crack

import (
	"fmt"
	"testing"

	"github.com/lightningnetwork/lnd/aezeed"
	"golang.org/x/exp/slices"
)

var fullMnemonic = []string{
	"above", "bronze", "era", "decade", "crane", "fossil", "hand", "tomato",
	"midnight", "entry", "bridge", "know", "fat", "power", "vintage", "uncle",
	"dinner", "carry", "train", "hollow", "parrot", "burger", "anchor", "clown",
}

func TestFindSingleMissingWord(t *testing.T) {
	for i := 0; i < len(fullMnemonic); i++ {
		i := i
		t.Run(fmt.Sprintf("missing-%d", i), func(t *testing.T) {
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
