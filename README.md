# dvm
dotnet core version manager

# Features
| Command | Description|
|-----| ----|
|current      |This command use to show you current dotnet core SDK version |
|  help       |Help about any command|
|  install    |This command use to install sdk|
|  list       |list you local installed dotnet core sdk version|
|  listAll    |This command to get all of the dotnet core sdk versions|
|  uninstall  |This command use to uninstall sdk|
|  upgrade    |Upgrade your dvm version|
|  use        |change local dotnet core version|
|  version    |This command to show the DVM version|


[![DEMO VIDEO](https://img.youtube.com/vi/2cLqHbYQ60I/0.jpg)](https://www.youtube.com/watch?v=2cLqHbYQ60I)



# How to use
`dvm help` you will get all the feature about the `dvm` 

# Install
please use [1.0.0's release](https://github.com/mangostano/dvm/tree/1.0.0)

> `curl https://raw.githubusercontent.com/mangostano/dvm/master/install.sh >> installDvm.sh && bash installDvm.sh`  

or   

> `wget https://raw.githubusercontent.com/mangostano/dvm/master/install.sh >> installDvm.sh && bash installDvm.sh`
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
