#!/bin/sh
set -e

REPO="dangerous-person/dopogoto"
BINARY="dopogoto"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  darwin)  PLATFORM="darwin_universal" ;;
  linux)   PLATFORM="linux_amd64" ;;
  *)       echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64) ;;  # already handled above
  arm64|aarch64) ;;  # darwin universal covers this
  *)
    if [ "$OS" = "linux" ]; then
      echo "Unsupported architecture: $ARCH (only amd64 is supported on Linux)"; exit 1
    fi
    ;;
esac

# Get latest release tag
echo "Finding latest release..."
TAG=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
  | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p')

if [ -z "$TAG" ]; then
  echo "Error: could not find latest release"; exit 1
fi

VERSION="${TAG#v}"
ARCHIVE="${BINARY}_${VERSION}_${PLATFORM}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${TAG}/${ARCHIVE}"

# Download and extract
TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

echo "Downloading ${BINARY} ${TAG}..."
curl -fsSL "$URL" -o "${TMPDIR}/${ARCHIVE}"
tar -xzf "${TMPDIR}/${ARCHIVE}" -C "$TMPDIR"

# Remove macOS quarantine attribute
if [ "$OS" = "darwin" ]; then
  xattr -d com.apple.quarantine "${TMPDIR}/${BINARY}" 2>/dev/null || true
fi

# Install binary
if [ -w /usr/local/bin ]; then
  INSTALL_DIR="/usr/local/bin"
else
  INSTALL_DIR="${HOME}/.local/bin"
  mkdir -p "$INSTALL_DIR"
fi

mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo ""
echo "Installed ${BINARY} ${TAG} to ${INSTALL_DIR}/${BINARY}"

if [ "$INSTALL_DIR" = "${HOME}/.local/bin" ]; then
  case ":$PATH:" in
    *":${INSTALL_DIR}:"*) ;;
    *) echo "Add ${INSTALL_DIR} to your PATH: export PATH=\"${INSTALL_DIR}:\$PATH\"" ;;
  esac
fi

echo "Run '${BINARY}' to start."
