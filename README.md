# dvm
dotnet core version manager

# Features



# How to use
`dvm help` you will get all the feature about the `dvm` 

# Install

> `curl https://raw.githubusercontent.com/mangostano/dvm/1.0.0/install.sh >> installDvm.sh && bash installDvm.sh`  

or   

> `wget https://raw.githubusercontent.com/mangostano/dvm/1.0.0/install.sh >> installDvm.sh && bash installDvm.sh`
## add the fellowing to you bash profile 
```
export DVM_HOME=$HOME/.dvm
export DOTNET_HOME=$HOME/.dotnet
export PATH=$DVM_HOME:$PATH
export PATH=$DOTNET_HOME:$PATH
```
(example) if you are using [zsh](https://github.com/robbyrussell/oh-my-zsh), add to your `.zshrc` file.

# Uninstall
remove the `.dvm` dictionary 

exam `sudo rm -f ~/.dvm`

# How to contribute 
the branch name is follow the [git flow](https://jeffkreeftmeijer.com/git-flow/) , the branch name is git flow default naming conversion

# [MIT LICENSE](./LICENSE)
