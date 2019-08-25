#!/bin/bash

while true
do
  age=$(($RANDOM % 10))
  if [ $(($RANDOM % 3)) -eq 0 ]; then
    name=bob
  else
    name=alice
  fi
  echo '{"age": '${age}', "name":"'${name}'"}'
done
