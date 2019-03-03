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

get_installed_dotnet(){
    echo dotnet --version
}

command_installed_path(){
  echo which $1
}

# check the curl or wget command
if [[ 0 -eq $(has_command curl) ]]; then
     DVM_INATSLL_COMMAND="curl"
elif [ 0 -eq $(has_command wget) ]; then
     DVM_INATSLL_COMMAND="wget"
  else
     echo "[ERROR] please install `wget` or `curl` command first"
     exit 1
fi

echo "START INSTALL DVM COMMAND"

# create the dvm home dictionary && dotnet command will installed in dvm root dictionary
mkdir -p ${DVM_HOME}/sdks
mkdir -p ${DVM_HOME}/scripts

if [[ ! -f "$DMV_HOME/$INSTALL_FILE_NAME" ]]; then
  rm -rf $INSTALL_FILE_NAME
fi

# user had installed dotnet by other way
if [[ 1 -eq $(has_command dotnet) ]]; then
    echo "Your had installed the dotnet by other ways, dvm is handle.\n After this, you can continue use the current version"
    sudo rm /etc/paths.d/dotnet
    CURRENT_DOTNET_VERSION = $(get_installed_dotnet)
    mv -f $(command_installed_path dotnet)/sdk $DVM_HOME/sdks
fi

$DVM_INSTALL_COMMAND https://dot.net/v1/dotnet-install.sh > $DVM_HOME/scripts/$INSTALL_FILE_NAME

chmod +X $DVM_HOME/scripts/$INSTALL_FILE_NAME

echo "# This is for DVM command" >> 

# get the current shell and remind the user to update the $PATH
echo "Please update the to"
echo "INSTALLATION COMPLETED, PLEASE ENJOY!"
# bash $INSTALL_FILE_NAME -Channel 1.0
# sudo ln -s $USERHOME/.dotnet/sdk/1.1.11 /usr/local/share/dotnet/sdk/
