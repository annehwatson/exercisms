#!/usr/bin/env bash

main () {
  result=""
  number=$1

  if (( number % 3 == 0))
  then
    result+="Pling"
  fi

  if (( number % 5 == 0))
    then
      result+="Plang"
  fi

  if (( number % 7 == 0))
    then
      result+="Plong"
  fi

  if [[ -z "$result" ]]
  then
    echo $number
  else
    echo $result
  fi
}

# call main with all of the positional arguments
main "$@"

