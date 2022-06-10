# passgen

A password generator that takes an easy to remember passphrase and turns it into a long, complex password.

### Usage

```
passgen [passphrase]

Description:
    generate secure passwords from a simple passphrase

Arguments:
    passphrase              input passphrase, optional

Options:
    -v, --version           override passgen algorithm version
    -s, --salt              salt appended to passphrase, default env[PASSGEN_SALT]
    -l, --length            password length, default 40
    -p, --print             print password instead of copying
        --no-specials       no special characters
        --no-uppers         no uppercase characters
        --no-numbers        no number characters
        --custom-specials   custom special character set
```

By default, passgen will copy your password to the clipboard. You can print to the terminal instead with the `--print` option.

### Setup
- Install with Go

  `go install github.com/jhotmann/passgen@latest`

- Install with Homebrew

  `TODO`

- Configure a salt (this will make your passwords unique to you even if you use a very basic input passphrase)
  - Generate a long random string using any method, including passgen itself: `passgen --no-specials -l 80 somerandomstringhere` >> `E9Fz0Fxs1JTND8E6HxJvcLjMJK6ZaH2DjzY8Pmot85mtLwX8Z1QaCCo6sUWrHlP7gavv9aBs3MQ9WPQ1`
  - Set the environment variable `PASSGEN_SALT` with your salt value.
    - On Unix-like, set it in your `.bashrc`, `.zshrc`, or equivalent file.
    - On Windows, go to System, Advanced System Settings, Environment Variables and set the variable in your user variables.

### Algorithms

- v1 - a salt is appended to your input passphrase, a skein hash is generated, the hash is base91 encoded, and a substring is returned. Compatable with [node-passgen-cli](https://github.com/jhotmann/node-passgen-cli)

- v2 - a salt is appended to your input passphrase, a sha256 hash is generated, the hash is converted to a BigInt, encoded to the supplied character set, and a substring is returned.

It is recommended to use v2 (the default) because it adds the ability to customize the available characters which is handy when websites only allow select special characters. It also uses a common hashing algorithm and should be portable to more languages.