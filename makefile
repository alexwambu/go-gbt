
---

## 📝 `Makefile`
```makefile
# GBTNetwork Makefile

BINARY := beth
BUILDDIR := build/bin
SRCDIR := ./cmd/beth
GO ?= go
GOFLAGS := -v

.PHONY: all beth clean

all: beth

beth:
	@echo ">>> Building $(BINARY) for GBTNetwork..."
	@mkdir -p $(BUILDDIR)
	$(GO) build $(GOFLAGS) -o $(BUILDDIR)/$(BINARY) $(SRCDIR)

clean:
	@echo ">>> Cleaning build artifacts..."
	@rm -rf $(BUILDDIR)
