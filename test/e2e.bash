#!/usr/bin/env bash

expected='{"a":7,"b":2}'
actual=$(echo 'a a a a b a a a b' | tr ' ' '\n' | scount -q)
if [[ $actual != $expected ]]; then
  echo "sample_data.txt: expected ${expected} but got ${actual}"
  exit 1
fi

echo "(test/e2e.bash) All test passed"
