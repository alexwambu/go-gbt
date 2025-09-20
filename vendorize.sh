#!/usr/bin/env bash
set -euo pipefail

# Require go.mod present
if [ ! -f go.mod ]; then
  echo "go.mod not found — create go.mod first"
  exit 1
fi

echo "Removing old vendor/ and go.sum..."
rm -rf vendor
rm -f go.sum

echo "Fetching modules..."
go mod tidy

echo "Vendoring..."
go mod vendor

echo "Verifying..."
go mod verify

echo "Done. You should now git add vendor go.sum go.mod and commit them."
echo "To commit:"
echo "  git add vendor go.sum go.mod"
echo "  git commit -m \"vendor: add go-ethereum v1.14.9 and deps\""
