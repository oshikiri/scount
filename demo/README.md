demo data for counting
=====

## Using generate_demo_jsonl.bash
```sh
./demo/generate_demo_jsonl.bash | jq -rc .age | scount
```

## Setup text8 corpus
```sh
wget http://mattmahoney.net/dc/text8.zip
unzip text8.zip
```

```sh
cat ./demo/text8 | tr ' ' '\n' | scount > /dev/null
```


```
$ cat ./demo/text8 | tr ' ' '\n' | grep -P '^(the|of|and|one|in|a|to|zero|nine|two|is)$' | sort | uniq -c | sort -r
1061396 the
 593677 of
 416629 and
 411764 one
 372201 in
 325873 a
 316376 to
 264975 zero
 250430 nine
 192644 two
 183153 is
```
