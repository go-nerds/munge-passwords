
![Munge passwords by aavision](https://github.com/aavision/munge-passwords/blob/main/favicon.png?raw=true)


# munge-passwords

A little mapped Go script to munge dictionary words into possible passwords.

## Usage


```bash
  go run .
```

```bash
Usage: munge-passwords.exe --level LEVEL [--wordlist WORDLIST] --output OUTPUT [--r] [--input INPUT]

Options:
  --level LEVEL, -l LEVEL
  --wordlist WORDLIST
  --output OUTPUT, -o OUTPUT
  --d, -d
  --input INPUT, -i INPUT
  --help, -h             display this help and exit
```

- `--level [1,9]` or `-l` : is the level of munge for the input or the word-list! 
- `--wordlist` or `-w` : The path of the wordlist.
- `--output` or `-o` : Output file name example `munge-passwords.txt`
- `--d` : Remove the duplicated record before munge from the provided list!
- `-i` or  `--input` : To enter string!

## License
