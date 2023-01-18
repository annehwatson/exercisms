#!/usr/bin/env bash

main () {
  strand_one=$1
  strand_two=$2
  hamming_distance=0

  if [ $# -ne 2 ]
  then
    echo "Usage: hamming.sh <string1> <string2>"
    return 1
  fi

  if [ ${#strand_one} -eq ${#strand_two} ]
  then
    for (( i=0; i<${#strand_one}; i++))
    do
      if [ "${strand_one:$i:1}" != "${strand_two:$i:1}" ]
      then
        hamming_distance=$((hamming_distance+1))
      fi
    done
  else
    echo "strands must be of equal length"
    return 1
  fi

  echo ${hamming_distance}
  return 0
}

main "$@"

