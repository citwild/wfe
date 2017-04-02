#!/usr/bin/env bash
set -e

GOFMT_PATHS=$(find . -not -wholename "*.git*" -not -wholename "*vendor*" -not -name "." -type d)

BAD_FILES=$(gofmt -s -l $GOFMT_PATHS)

if [ -n "$BAD_FILES" ]; then
  echo "The following files are not properly formatted:"
  echo $BAD_FILES
  exit 1
fi