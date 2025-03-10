#!/bin/sh

BIN_DIR=$(pwd)

THE_ARCH_BIN=''
THIS_PROJECT_NAME='wdhd'
OWNER='kevincobain2000'

THISOS=$(uname -s)
ARCH=$(uname -m)

INSTALL_VERSION=${INSTALL_VERSION:-latest}
echo "Installing $THIS_PROJECT_NAME version: $INSTALL_VERSION"

case $THISOS in
   Linux*)
      case $ARCH in
        arm64)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-linux-arm64"
          ;;
        aarch64)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-linux-arm64"
          ;;
        armv6l)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-linux-arm"
          ;;
        armv7l)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-linux-arm"
          ;;
        *)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-linux-amd64"
          ;;
      esac
      ;;
   Darwin*)
      case $ARCH in
        arm64)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-darwin-arm64"
          ;;
        *)
          THE_ARCH_BIN="$THIS_PROJECT_NAME-darwin-amd64"
          ;;
      esac
      ;;
   Windows|MINGW64_NT*)
      THE_ARCH_BIN="$THIS_PROJECT_NAME-windows-amd64.exe"
      THIS_PROJECT_NAME="$THIS_PROJECT_NAME.exe"
      ;;
esac

if [ -z "$THE_ARCH_BIN" ]; then
   echo "This script is not supported on $THISOS and $ARCH"
   exit 1
fi

DOWNLOAD_URL="https://github.com/$OWNER/$THIS_PROJECT_NAME/releases/download/$INSTALL_VERSION/$THE_ARCH_BIN"
if [ "$INSTALL_VERSION" = "latest" ]; then
  DOWNLOAD_URL="https://github.com/$OWNER/$THIS_PROJECT_NAME/releases/$INSTALL_VERSION/download/$THE_ARCH_BIN"
fi

echo "Downloading from $DOWNLOAD_URL..."
HTTP_STATUS=$(curl -kL --progress-bar -w "%{http_code}" -o "$BIN_DIR/$THIS_PROJECT_NAME" "$DOWNLOAD_URL")

if [ "$HTTP_STATUS" -ne 200 ]; then
  echo "Error: Failed to download $THIS_PROJECT_NAME. HTTP status code: $HTTP_STATUS"
  exit 1
fi

chmod +x "$BIN_DIR/$THIS_PROJECT_NAME"

echo "Installed successfully to: $BIN_DIR/$THIS_PROJECT_NAME"
