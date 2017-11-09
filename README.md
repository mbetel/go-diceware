## Golang Diceware Generator

This library implements the [Diceware](https://en.wikipedia.org/wiki/Diceware)
algorithm in pure Golang. The algorithm is most-commonly used when generating
human-readable passwords.

The list of words are generated from the [EFF's "long"
list](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).
However, the API's are abstracted so you can roll die and then use your own word
list as-needed.

It uses crypto/rand for rolling die for added randomness.

## Installation

```sh
$ go -u get github.com/sethvargo/diceware
```

## Usage

```golang
package main

import (
  "strings"

  "github.com/sethvargo/go-diceware/diceware"
)

func main() {
  // Generate 6 words using the diceware algorithm.
  list, err := diceware.Generate(6)
  if err != nil  {
    log.Fatal(err)
  }
  log.Printf(strings.Join(list, "-"))
}
```

## License

This code is licensed under the MIT license.