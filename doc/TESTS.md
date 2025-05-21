# Tests

This document describes a manual test protocol.

## info splash screen

+ info splash screen is displayed initially in **normal mode** (no flags)
+ info splash screen is displayed initially in **dev mode** (`--dev` flag)
+ info splash screen is skipped in **consent mode** (`--consent` flag)
+ quit **exits program**
+ **default selection** is quit
+ quit is **disabled until seen**, i.e. scrolled to bottom
+ **ok closes** info splash screen
+ **ok navigates to menu**

## main menu

### navigation

+ menu can be navigated with **up/down arrow keys**
+ menu can be navigated with **TAB and BACKTAB**
+ **CTRL+P** should move cursor to previous line
+ **CTRL+N** should move cursor to next line
+ **SPACE** in main menu should have no effect (reserved for future options)
+ **ENTER** selects menu item
+ **ESC** offers to quit application
+ **no other keys** have effect
+ navigation **wraps around**
+ **footer hints** are displayed when navigating
+ menu items can be selected with **shortcuts**
+ **no selection with cursor keys** of disabled menu items
+ **no selection with shortcuts** of disabled menu items
+ **disabled menu items are skipped** when navigating

### item visibility

+ random, display, swapped are only **visible in dev mode** (`--dev` flag)
+ navigation back to info is **visible in normal mode** (no flags)
+ navigation back to info is **visible in consent mode** (`--consent`)
+ only finding wrong or swapped word are enabled when mnemonic is **invalid**
+ only finding missing word is enabled when mnemonic is **incomplete**
+ all brute force options are disabled when mnemonic is **valid**

### info item

### random mnemonic (dev)

+ **custom password** is honoured in seed generation

### edit mnemonic (see seed view)

### clear mnemonic

### find missing word

+ finding missing word works for **custom password**
+ when missing word is found, the **seed view is updated** accordingly
+ **highlight** (background) the missing words in seed view

#### edge cases

+ works for **last word**
+ works for **first word**

### find swapped word (dev)

+ finding swapped word works for **custom password**
+ when swapped word is found, the **seed view is updated** accordingly
+ **highlight** (background) the swapped words in seed view

#### edge cases

+ works for **last word**
+ works for **first word**

### find wrong word

+ finding wrong word works for **custom password**
+ when wrong word is found, the **seed view is updated** accordingly
+ **highlight** (background) the replaced words in seed view
+ **irretrievable mnemonic** (wrong password, too many wrong words) returns at
  some
  point with failure

#### edge cases

+ works for **last word**
+ works for **first word**

### display seed (dev)

+ **navigates to result** view

### quit

## seed view

+ password and mnemonic are properly **aligned**
+ omitting more than the **amount of recoverable words** does not activate the
  missing words feature

### navigation

+ **CTRL-N** navigates down
+ **CTRL-P** navigates up
+ **CTRL-A** should move cursor to beginning of line
+ **CTRL-E** should move cursor to end of line
+ **CTRL-DEL** deletes word
+ navigation **wraps around**

### word input fields

+ **footer hint is displayed** when navigating into word input field
+ leaving word input field **removes footer hint** (even if there is no menu
  item hint)
+ **SPACE** on a seed word should activate word selection, even on complete and
  valid word
+ **pasting mnemonic** from clipboard populates words
+ valid/invalid words are **visually indicated**
+ **ESC with open dropdown** in word input does not update mnemonic but restores
  the previous word
+ **ESC when navigating** should return to main menu
+ **ENTER** moves on to next menu item

### password input field

+ **footer hint is displayed** when navigating into password input field
+ leaving password input field **removes footer hint** (even if there is no menu
  item hint)
+ **pasting is disabled** in password input field
+ **ESC** returns to main menu
+ **ENTER** moves on to first word input

### validation view

+ initially displaying '**incomplete mnemonic**' when mnemonic is incomplete
+ displaying '**invalid mnemonic**' when mnemonic is complete but invalid
+ displaying '**invalid words**' when at least one word is not in wordlist
+ displaying '**valid mnemonic**' when password and mnemonic are valid

## result view

+ QR-code is **displayed**
+ QR-Code can be **scanned in light theme**
+ QR-Code can be **scanned in dark theme**
+ **link caption** is present
+ **ok button** leaves application

## footer

+ **progress** is displayed
+ footer text for **successful recovery** contains (tries, time, etc.)
+ footer text for **ongoing brute force attempt** contains (tries, time, etc.)
+ footer text for **failed brute force attempt** contains (tries, time, etc.)

## resizing

+ resizing window does **not break layout**
+ it is not possible to **scroll out of bounds**
+ **no menu items are hidden** when resizing, unless strictly necessary

## documentation

### [README.md](./README.md)

+ documentation is orthographically correct
+ functionality is described
+ known issues are described
+ author is referenced
+ license is referenced
