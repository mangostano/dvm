#!/bin/sh
INSTALL_FILE_NAME="install.sh"
DVM_HOME="$HOME/.dvm"
DOTNET_HOME="$HOME/.dotnet"
DVM_COMMAND_REPO="https://raw.githubusercontent.com/mangostano/dvm/1.0.0/dvm"

has_command() {
  if ! command -v "$1" > /dev/null 2>&1
    then echo 1;
  else
    echo 0;
  fi
}

get_dotnet_version(){
    dotnet --version
}

get_dotnet_path(){
    INSTALLED_DOTNET_PATH=$(command -v dotnet)
    echo "${INSTALLED_DOTNET_PATH%dotnet}"
}

get_user_name(){
   whoami
}

if [[ 0 -eq $(has_command dvm) ]]; then
     echo "[INFO] dvm command already installed"
     exit 1;
fi

# check the curl or wget command
if [[ 0 -eq $(has_command curl) ]]; then
     DOWNLOAD_COMMAND="curl"
elif [[ 0 -eq $(has_command wget) ]]; then
     DOWNLOAD_COMMAND="wget"
  else
     echo "[ERROR] please install `wget` or `curl` command first"
     exit 1
fi

echo "START INSTALL DVM COMMAND"

# create the dvm home dictionary && dotnet command will installed in dvm root dictionary
mkdir -p ${DVM_HOME}/sdks
mkdir -p ${DVM_HOME}/scripts
mkdir -p ${DOTNET_HOME}

if [[ ! -f "${DMV_HOME}/scripts/$INSTALL_FILE_NAME" ]]; then
  rm -f "${DMV_HOME}/scripts/$INSTALL_FILE_NAME"
fi

# user had installed dotnet by other way
if [[ 0 -eq $(has_command dotnet) ]]; then
    printf "Your had installed the dotnet by other ways, dvm is handle.\n After this, you can continue use the current version\n"
    sudo rm /etc/paths.d/dotnet
    DOTNET_VERSION=$(get_dotnet_version)
    # mv the installed sdk to dvm/SDKs
    sudo mv -f $(get_dotnet_path)/sdk/* ${DVM_HOME}/sdks
    sudo mv -f $(get_dotnet_path)/* ${DOTNET_HOME}/
    sudo chown -R $(get_user_name) ${DVM_HOME} ${DOTNET_HOME}
    rm -rf ${DVM_HOME}/sdks/N*
    ln -s ${DVM_HOME}/sdks/${DOTNET_VERSION} ${DOTNET_HOME}/sdk/
fi

# curl the microsoft dotnet install script
${DOWNLOAD_COMMAND} https://dot.net/v1/dotnet-install.sh > ${DVM_HOME}/scripts/${INSTALL_FILE_NAME} && chmod +x ${DVM_HOME}/scripts/${INSTALL_FILE_NAME}

# curl the dvm command from repo && need user to update path
${DOWNLOAD_COMMAND} ${DVM_COMMAND_REPO} > ${DVM_HOME}/dvm && chmod +x ${DVM_HOME}/dvm

printf "\n# This is for DVM command\n"
printf "Please add the\n\n"
printf "export DVM_HOME=$DVM_HOME\n"
printf "export DOTNET_HOME=$HOME/.dotnet\n"
printf "export PATH=\$DOTNET_HOME:\$PATH\n"
printf "export PATH=\$DVM_HOME:\$PATH\n\n"
printf "to your PATH variable\n"
printf "INSTALLATION COMPLETED, PLEASE ENJOY!"
