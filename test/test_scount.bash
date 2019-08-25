#!/usr/bin/env bash

expected='{"a":7,"b":2}'
actual=$(cat ./demo/sample_data.txt | scount)
if [[ $actual != $expected ]]; then
  echo "sample_data.txt: expected ${expected} but got ${actual}"
  exit 1
fi

echo "(test_scount.bash) All test passed"
