#!/bin/bash

function main() {
    test -f main.go || return

    go build main.go

    for path in tests/*; do
      [[ "${path}" =~ a$ ]] && continue

      echo -n "${path}: "
      cat "${path}" | ./main > my_output
      diff "${path}.a" my_output > /dev/null && echo "OK" || echo "FAIL"
    done

    rm main my_output
}

main
