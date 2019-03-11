#!/bin/sh
INSTALL_FILE_NAME="install.sh"
DVM_HOME="$HOME/.dvm"
DVM_COMMAND_REPO="https://github.com/mangostano/dvm/blob/master/dvm"

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

installed_dotnet_path(){
    INSTALLED_DOTNET_PATH=$(command -v dotnet)
    echo "${INSTALLED_DOTNET_PATH%dotnet}"
}

# check the curl or wget command
if [[ 0 -eq $(has_command curl) ]]; then
     DVM_INSTALL_COMMAND="curl"
elif [ 0 -eq $(has_command wget) ]; then
     DVM_INSTALL_COMMAND="wget"
  else
     echo "[ERROR] please install `wget` or `curl` command first"
     exit 1
fi

echo "START INSTALL DVM COMMAND"

# create the dvm home dictionary && dotnet command will installed in dvm root dictionary
mkdir -p ${DVM_HOME}/sdks
mkdir -p ${DVM_HOME}/scripts

if [[ ! -f "$DMV_HOME/scripts/$INSTALL_FILE_NAME" ]]; then
  rm -f "$DMV_HOME/scripts/$INSTALL_FILE_NAME"
fi

# user had installed dotnet by other way
if [[ 0 -eq $(has_command dotnet) ]]; then
    printf "Your had installed the dotnet by other ways, dvm is handle.\n After this, you can continue use the current version"
    sudo rm /etc/paths.d/dotnet
    CURRENT_DOTNET_VERSION=$(get_installed_dotnet)
    # mv the installed sdk to dvm/SDKs
    sudo mv -rf $(installed_dotnet_path)/sdk/* $DVM_HOME/sdks
fi

# curl the microsoft dotnet install script
$DVM_INSTALL_COMMAND https://dot.net/v1/dotnet-install.sh > $DVM_HOME/scripts/$INSTALL_FILE_NAME && chmod +X $DVM_HOME/scripts/$INSTALL_FILE_NAME

# curl the dvm command from repo && need user to update path
$DVM_INSTALL_COMMAND $DVM_COMMAND_REPO > $DVM_HOME/dvm && chomd +x $DVM_HOME/dvm

printf "# This is for DVM command\n"
printf "Please add the\n\n"
printf "export DVM_HOME=$DVM_HOME\n"
printf "export PATH=\$HOME/.dvm:\$PATH\n\n"
printf "to your PATH variable\n"
printf "INSTALLATION COMPLETED, PLEASE ENJOY!"
