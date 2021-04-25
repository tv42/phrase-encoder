# eagain.net/go/phrase-encoder -- Encode binary data as words

```
go install eagain.net/go/phrase-encoder/cmd/phrase-encode
go install eagain.net/go/phrase-encoder/cmd/phrase-decode
```

These command-line utilities wrap the library https://gitlab.com/NebulousLabs/entropy-mnemonics and use it to encode and decode data to/from words.

```
$ echo test | phrase-encode
lazy-snout-joining-aces
$ echo lazy-snout-joining-aces | phrase-decode
test
```

This can be also be used to create easy-to-type random passwords:

```
$ dd if=/dev/urandom bs=16 count=1 status=none | phrase-encode
tamper-dogs-solved-jeers-gambit-utensils-simplest-yields-drunk-tulips-mews-king
```

(See also [`eagain.net/go/entropy`](https://github.com/tv42/entropy).)

The output words are joined with dashes for easier copy-pasting.

To make the encoded data easier to type, input can use spaces instead of dashes:

```
$ phrase-decode
lazy snout joining aces
^D
test
```
