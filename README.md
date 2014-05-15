# quintessence - simple golang gist uploader

> quin·tes·sence  _(noun)_
- the aspect of something regarded as the intrinsic and central constituent of its character.
- a refined essence or extract of a substance.

## What

`quint` reads text from STDIN and publishes it as an anonymous [gist](https://gist.github.com/).

# How
```sh
$ cat main.go | quint
https://gist.github.com/2592adc620f2c019b1ea

$ quint main.go 
https://gist.github.com/2cc21997dee4d12a65e9
```

## Where

[Binary releases are available](https://github.com/dpritchett/quintessence/releases) for Linux, OSX, and Windows.
