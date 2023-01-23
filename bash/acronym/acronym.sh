#!/usr/bin/env bash

# loop over each word and append first character to result
# what about hyphens?

main () {
  IFS='-_* '
  words=$1
  result=()
  for word in $words
  do
    uppercase_word="${word^^}"
    result+=${uppercase_word:0:1}
  done
  echo ${result}
}

main "$@"

