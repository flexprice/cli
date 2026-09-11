#!/bin/sh

set -eu

BINARY="flexprice"
BASE_URL="https://github.com/flexprice/cli/releases"

if [ -t 2 ] && [ -z "${NO_COLOR:-}" ]; then
  BOLD="$(printf '\033[1m')"
  GREEN="$(printf '\033[32m')"
  YELLOW="$(printf '\033[33m')"
  RED="$(printf '\033[31m')"
  RESET="$(printf '\033[0m')"
else
  BOLD="" GREEN="" YELLOW="" RED="" RESET=""
fi

info() { printf '%s\n' "$*" >&2; }
ok()   { printf '%s✓%s %s\n' "$GREEN" "$RESET" "$*" >&2; }
warn() { printf '%s!%s %s\n' "$YELLOW" "$RESET" "$*" >&2; }
fail() { printf '%serror:%s %s\n' "$RED" "$RESET" "$*" >&2; exit 1; }

need_cmd() {
  command -v "$1" >/dev/null 2>&1 || fail "required command not found: $1"
}

detect_os() {
  case "$(uname -s)" in
    Darwin) echo darwin ;;
    Linux)  echo linux ;;
    MINGW*|MSYS*|CYGWIN*|Windows_NT)
      fail "Windows is not supported by this script. Download the .zip from ${BASE_URL}/latest or run: go install github.com/flexprice/cli/cmd/flexprice@latest" ;;
    *) fail "unsupported operating system: $(uname -s)" ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64)   echo amd64 ;;
    arm64|aarch64)  echo arm64 ;;
    *) fail "unsupported architecture: $(uname -m) (releases are built for amd64 and arm64)" ;;
  esac
}

download() {
  # download <url> <dest>
  if command -v curl >/dev/null 2>&1; then
    curl -fsSL --retry 3 --retry-delay 1 -o "$2" "$1"
  elif command -v wget >/dev/null 2>&1; then
    wget -q -O "$2" "$1"
  else
    fail "need curl or wget to download releases"
  fi
}

resolve_latest_version() {
  # Resolve the tag via redirect to avoid GitHub API rate limits.
  if command -v curl >/dev/null 2>&1; then
    curl -fsSLI -o /dev/null -w '%{url_effective}' "${BASE_URL}/latest" \
      | sed 's|.*/tag/||'
  elif command -v wget >/dev/null 2>&1; then
    wget -q -S --spider "${BASE_URL}/latest" 2>&1 \
      | grep -i '^ *Location:' | tail -n 1 | sed 's|.*/tag/||' | tr -d '\r'
  else
    fail "need curl or wget to resolve the latest version"
  fi
}

sha256_of() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$1" | awk '{print $1}'
  elif command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1" | awk '{print $1}'
  else
    fail "need sha256sum or shasum to verify the download"
  fi
}

choose_install_dir() {
  if [ -n "${FLEXPRICE_INSTALL_DIR:-}" ]; then
    echo "$FLEXPRICE_INSTALL_DIR"
  elif [ -d /usr/local/bin ] && [ -w /usr/local/bin ]; then
    echo /usr/local/bin
  else
    echo "${HOME}/.local/bin"
  fi
}

on_path() {
  case ":${PATH}:" in
    *":$1:"*) return 0 ;;
    *) return 1 ;;
  esac
}

main() {
  need_cmd uname
  need_cmd tar
  need_cmd mktemp

  OS="$(detect_os)"
  ARCH="$(detect_arch)"

  VERSION="${FLEXPRICE_VERSION:-}"
  if [ -z "$VERSION" ]; then
    VERSION="$(resolve_latest_version)"
    [ -n "$VERSION" ] || fail "could not determine the latest release; set FLEXPRICE_VERSION=vX.Y.Z"
  fi
  case "$VERSION" in
    v*) ;;
    *) VERSION="v${VERSION}" ;;
  esac
  BARE_VERSION="${VERSION#v}"

  ARCHIVE="${BINARY}_${BARE_VERSION}_${OS}_${ARCH}.tar.gz"
  ARCHIVE_URL="${BASE_URL}/download/${VERSION}/${ARCHIVE}"
  CHECKSUMS_URL="${BASE_URL}/download/${VERSION}/checksums.txt"

  INSTALL_DIR="$(choose_install_dir)"

  info "${BOLD}Installing ${BINARY} ${VERSION} (${OS}/${ARCH})${RESET}"

  TMP_DIR="$(mktemp -d 2>/dev/null || mktemp -d -t flexprice)"
  trap 'rm -rf "$TMP_DIR"' EXIT INT TERM

  info "Downloading ${ARCHIVE_URL}"
  download "$ARCHIVE_URL" "${TMP_DIR}/${ARCHIVE}" \
    || fail "download failed. Check that ${VERSION} exists at ${BASE_URL}"
  download "$CHECKSUMS_URL" "${TMP_DIR}/checksums.txt" \
    || fail "could not download checksums.txt for ${VERSION}"

  EXPECTED="$(awk -v f="$ARCHIVE" '$2 == f {print $1}' "${TMP_DIR}/checksums.txt")"
  [ -n "$EXPECTED" ] || fail "no checksum entry for ${ARCHIVE} in checksums.txt"
  ACTUAL="$(sha256_of "${TMP_DIR}/${ARCHIVE}")"
  if [ "$EXPECTED" != "$ACTUAL" ]; then
    fail "checksum mismatch for ${ARCHIVE}
  expected: ${EXPECTED}
  actual:   ${ACTUAL}
The download may be corrupted or tampered with. Nothing was installed."
  fi
  ok "Checksum verified"

  tar -xzf "${TMP_DIR}/${ARCHIVE}" -C "$TMP_DIR" "$BINARY" \
    || fail "could not extract ${BINARY} from ${ARCHIVE}"

  mkdir -p "$INSTALL_DIR" || fail "cannot create ${INSTALL_DIR}"
  [ -w "$INSTALL_DIR" ] || fail "${INSTALL_DIR} is not writable. Set FLEXPRICE_INSTALL_DIR to a directory you own, or re-run with sudo."

  # Rename atomically to avoid overwriting a running binary.
  install_tmp="${INSTALL_DIR}/.${BINARY}.tmp.$$"
  cp "${TMP_DIR}/${BINARY}" "$install_tmp"
  chmod 0755 "$install_tmp"
  mv -f "$install_tmp" "${INSTALL_DIR}/${BINARY}"

  ok "Installed ${INSTALL_DIR}/${BINARY}"

  if on_path "$INSTALL_DIR"; then
    info ""
    info "Get started with:"
    info "  ${BINARY} init"
  else
    warn "${INSTALL_DIR} is not on your PATH. Add it with:"
    info "  echo 'export PATH=\"${INSTALL_DIR}:\$PATH\"' >> ~/.zshrc   # or ~/.bashrc"
    info "  source ~/.zshrc"
    info ""
    info "Then get started with:"
    info "  ${BINARY} init"
  fi
}

main "$@"
