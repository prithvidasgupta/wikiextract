#!/usr/bin/env bash

package=$1

if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi