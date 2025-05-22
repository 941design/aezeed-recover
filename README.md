# aezeed-recover

Recover wrong or missing words from an aezeed mnemonic.

## Overview

_aezeed-recover_ does not need an internet connection, and works well on a
80x24 standard terminal, i.e. on your node or a primitive device for the very
purpose of recovery.

<img alt="demo" src="/doc/demo.gif?raw=true" width="520px"/>

+ Currently, recovery is limited to 2 missing or a single wrong word
+ Replacements and insertions are tried at any position
+ Password recovery is not supported, yet

(see [TODO](#todo) for upcoming features)

## Similar Projects

+ [btcrecover](https://github.com/gurnec/btcrecover/) written in python
  by [gurnec](https://github.com/gurnec)
+ [bip39-solver-cpu](https://github.com/johncantrell97/bip39-solver-cpu) written
  in rust by [johncantrell97](https://github.com/johncantrell97)
+ [bip39-solver-gpu](https://github.com/dessy888/bip39-solver-gpu) written in
  rust by [dessy888](https://github.com/dessy888)
+ [Seed Savior](https://github.com/ZenGo-X/mnemonic-recovery) written in
  javascript by [ZenGo-X](https://github.com/ZenGo-X)
+ [walletsrecovery.org](https://walletsrecovery.org/) overview of
  custom derivation paths of different wallet implementations

## Installation

Download the appropriate binary
from [releases](https://github.com/941design/aezeed-recover/releases), or build
from source:

```sh
git clone https://github.com/941design/aezeed-recover.git
cd aezeed-recover
go install
```

## Running

Either run the binary directly, or execute `aezeed-recover` if installed
with `go install`. Make sure `~/go/bin/` is on your `$PATH`.

### Dev Mode

Running with the `--dev` flag will enable random seed generation and other
options.

DO NOT USE THE GENERATED MNEMONIC FOR REAL FUNDS.

## Known Issues

+ skipping through seed words with arrow keys occasionally activates drop down
  unintentionally
+ when skipping through seed words with arrows, the cursor is not always
  positioned at the end of the word
+ on the info splash screen the confirmation button is disabled initially,
  although the entire text is visible
+ pasting should NOT wrap around into password field
+ pasting non-alpha characters should not be possible
+ when leaving password field with ESC, the seed is not reevaluated
+ pasting from clipboard on linux and windows is untested

## TODO

### improve brute force by limiting realm

The amount of missing/wrong words can be increased significantly by considering
more detailed knowledge of the seed. If parts of the seed are definite, the
search space is vastly reduced. Possible improvements:

+ mark certain words as fixed
+ allow prefixes as given
+ allow suffixes as given
+ consider placeholders or a regex like DSL

See [btcrecover](https://github.com/gurnec/btcrecover/).

### improve brute force by trying more likely mistakes first

The following assumptions should be honoured:

+ adjacent words ar more likely to be mistaken than distant words
+ mistakes are made more likely towards the end of the mnemonic

### BIP39 support

Also recovering BIP39 seeds is desirable.

BIP39 only has a 4 bit checksum, as opposed to aezeed's 4 bytes. Hence, brute
forcing missing or wrong words requires knowledge of an address, its derivation
path, as well as access to a transaction index (a node).

### find swapped words and columns

Another possible cause of error is swapping or rotating the mnemonic when
writing it down, i.e. columns are mistaken for rows.

+ check for swaps along all symmetrical rotations, i.e. swap words between
  columns if a column contains 2,3,4,6,8,12 words

## Donation

BTC: bc1qfkv06pjcgq280f42nu66ktdfeer7f06lmtgerp

Thank you very much for your support!

## License

[MIT](./LICENSE.md)

## Author

[Markus Rother](markus.rother@mustwork.de)

---
