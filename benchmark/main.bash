#!/bin/bash

n_trials=10

git_version=$(git log -1 --pretty=%H)
output_path=results/benchmark_results.tsv

function run () {
  target_command=$1
  for i in $(seq $n_trials); do
    /usr/bin/time \
      -f "$target_command\t$i\t%e\t%M" \
      -o $output_path \
      -a cat ../demo/text8_0 | eval $target_command > /dev/null 2>&1
  done
}

echo 'command	i_trials	elapsed_time	memory_usage_kb' > $output_path

run 'scount -q'
run 'scount -q -a'
run 'scount'
run 'scount -a'

echo $git_version > ./results/environments.log
