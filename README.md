# PassGen (In Development)

A password generator that takes an easy to remember passphrase and turns it into a long, complex password.

### Usage

```
passgen [passphrase]

Description:
    generate secure passwords from a simple passphrase

Arguments:
    passphrase              input passphrase, optional

Options:
    -v, --version           override PassGen version
    -s, --salt              salt appended to passphrase, default env[PASSGEN_SALT]
    -l, --length            password length, default 40
    -p, --print             print password instead of copying
        --no-specials       no special characters
        --no-uppers         no uppercase characters
        --no-numbers        no number characters
        --custom-specials   custom special character set
```

### Setup
- Install with Go

  `go install github.com/jhotmann/passgen`

- Install with Homebrew

  `TODO`

### Algorithms

- v1 - a salt is appended to your input passphrase, a skein hash is generated, the hash is base91 encoded, and a substring is returned.

- v2 - a salt is appended to your input passphrase, a sha256 hash is generated, the hash is encoded to the supplied character set, and a substring is returned.

It is recommended to use v2 (the default) because it adds the ability to customize the available characters which is handy when websites only allow select special characters.