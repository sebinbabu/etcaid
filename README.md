# etcaid

`etcaid` manages your Linux application config backups with ease. It copies your application configuration to a centralized backup folder which can then be synced to multiple devices using third party tools (such as `git`).

## Quickstart

You must have Go installed on your system before continuing.

* Clone this repository.
* Run `make install` to install `etcaid` on your system.
```console
$ etcaid init # initializes the etcaid config & directories

$ etcaid new zsh # creates a new application config with the name zsh and opens in your editor (default=`vim`) for you
```
* `zsh.toml` opened in editor:
```toml
title = "zsh"
home_paths = [] # home_paths contains paths to files relative to the user's home path (eg: /home/currentuser)
xdg_config_paths = [] # xdg_config_path contains paths to files relative to the xdg application configuration directory for the user (eg: /home/currentuser/.config)
```
* You can edit this config to add all configuration file/directory paths for zsh.
```toml
title = "zsh"
home_paths = [
	".zshrc"
]
xdg_config_paths = []
```
* Once you've saved this and closed the editor:
```console
Application config for zsh edited at /home/currentuser/etcaid/applications/zsh.toml

$ etcaid backup # backups all application config files

copied /home/currentuser/.zshrc to /home/currentuser/etcaid/backup/zsh/home/.zshrc
```
* To restore:
```console
$ etcaid restore # restores all application config files

copied /home/currentuser/etcaid/backup/zsh/home/.zshrc to /home/currentuser/.zshrc
```
* Use `git` on the `~/etcaid` directory to save your backups and sync it across devices.

## Why `etcaid`?

While there are mature & amazing projects out there that would solve most problems of syncing application configurations across devices, I needed a tool that would be simple enough to just let me define what files needed to be backed up for an application. And I'd just go ahead and use `git` for versioning and syncing. No magic, no surprises. And thus, `etcaid` was born.

## Todo

1. Finish with tests
2. Take a backup of destination before replacing it while running `etcaid restore`.
