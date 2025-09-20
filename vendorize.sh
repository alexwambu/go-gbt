#!/bin/bash
set -e

echo "🚀 Starting vendoring process for GBTNetwork..."

# Ensure go.mod exists
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod not found. Run 'go mod init github.com/alexwambu/go-gbt' first."
    exit 1
fi

# Download all modules
echo "📦 Downloading dependencies..."
go mod download

# Verify modules
echo "🔍 Verifying module integrity..."
go mod verify

# Vendor modules into /vendor
echo "📂 Creating vendor directory..."
go mod vendor

# Tidy up go.mod and go.sum (optional, since you want static deployment)
echo "🧹 Tidying go.mod and go.sum..."
go mod tidy

echo "✅ Vendoring complete! Dependencies are now stored in ./vendor/"
echo "👉 You can now build with: go build -mod=vendor ./..."
