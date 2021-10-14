# dotsync
Dotsync is a small utility for easily symlinking "dot profiles" from the git repository to the homefolder

Dotsync can administer multiple dotprofiles by creating more folder in the `.dotfiles` folder in your home directory

## How to use dotsync

* You first need either clone this repository to your computer or download the script

* Create a folder called `.dotfiles` in your home folder
```bash
mkdir $HOME/.dotfiles
```

* Create a new profile in the `.dotfiles` folder (Profiles can not start with a "." dot)
```bash
mkdir -p $HOME/.dotfiles/bashprofile
```

* When you have a new profile, either copy your current dotfile into the profile or make new dotfiles

* Run dotsync
```bash
python3 dotsync
```

* choose the profile you want to install and just let it work it's magic