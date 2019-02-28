#!/bin/sh
INSTALL_FILE_NAME="install.sh"
DVM_HOME="$HOME/.dvm"


has_command() {
  if ! command -v "$1" > /dev/null 2>&1
    then echo 1;
  else 
    echo 0;
  fi
}

if [ 0 -eq $(has_command curl) ]; then
     DVM_INATSLL_COMMAND="curl"
elif [ 0 -eq $(has_command wget) ]; then
     DVM_INATSLL_COMMAND="wget"
  else 
     echo "[ERROR] please install `wget` or `curl` command first"
     exit 1
fi

echo "START INSTALL DVM COMMAND"
echo

mkdir -p ${DVM_HOME}/sdks

if [[ ! -f "$DMV_HOME/$INSTALL_FILE_NAME" ]]; then
  rm -rf $INSTALL_FILE_NAME
fi


$DVM_INATSLL_COMMAND https://dot.net/v1/dotnet-install.sh > $DVM_HOME/$INSTALL_FILE_NAME

sudo chmod +X $DVM_HOME/$INSTALL_FILE_NAME

echo "INSTALLATION COMPLETED, PELEASE ENJOY!"
# bash $INSTALL_FILE_NAME -Channel 1.0
# sudo ln -s $USERHOME/.dotnet/sdk/1.1.11 /usr/local/share/dotnet/sdk/
