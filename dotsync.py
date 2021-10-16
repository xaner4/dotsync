#!/usr/bin/env python3

import os
import shutil

"""
Dotsync
Created By: Marius (https://github.com/xaner4/)
License: MIT
"""

HOME = os.getenv('HOME')
INSTALL_DIR = os.path.join(HOME, '')
DOTDIR = os.path.join(HOME, ".dotfiles")
BACKUP_DIR = os.path.join(DOTDIR, ".backup")
HOSTNAME = os.getenv('HOSTNAME')

# TODO: Write better error messages
# TODO: Write better comments about the code
# TODO: Write better instructions


def get_profiles():
    return [d for d in os.listdir(DOTDIR) if os.path.isdir(os.path.join(DOTDIR, d)) and not d.startswith('.')]


def choose_profiles(p: list):
    for i, v in enumerate(p):
        # TODO: Find out why short profiles are not indented correctly
        print(f" [{i}]\t {v}:\t\t {DOTDIR}/{v}")
    ch = input("What profiles do you want to install: ")
    ch = list(map(int, ch.split(",")))
    return ch

def profiles_to_install(pr, ch):
    install = []
    for i in ch:
        t = pr[i]
        install.append(t)
    return install

def install(profile):
    profile_path = os.path.join(DOTDIR, profile)
    files = os.listdir(profile_path)
    for i, v in enumerate(files):
        src = os.path.join(profile_path, v)
        dst = os.path.join(INSTALL_DIR, v)

        if os.path.islink(dst):
            # TODO: Ask if want to overwrite exsisting link
            realpath = os.path.realpath(dst)
            print(f"{dst} is already a symlink to {realpath}")
            if realpath == src:
                continue

            answer = input(f"Do you want to overwrite link {dst} [yes/NO]: ").lower()
            if answer == "yes":
                os.unlink(dst)
            else:
                continue

        if os.path.exists(dst):
            backup(dst)
            try:
                delete(dst)
            except Exception as e:
                print(e)

        try:
            os.symlink(src, dst)
        except Exception as e:
            print(e)

        if not os.path.islink(dst):
            print(f"Something went wrong with linking {src} to {dst}")


def backup(dst):
    # TODO: Check if backup directory already exsisting 
    bd = os.path.join(BACKUP_DIR, HOSTNAME)
    if not os.path.isdir(bd):
        try:
            os.makedirs(bd)
        except Exception as e:
            print(e)
    if os.path.isdir(dst):
        backup_dir(dst, bd)
    else:
        backup_file(dst, bd)

def backup_dir(src, dst):
    base = os.path.basename(src)
    try:
        shutil.copytree(src, os.path.join(dst, base))
    except Exception as e:
        print(e)


def backup_file(src, dst):
    try:
        shutil.copy(src, dst)
    except Exception as e:
        print(e)

def delete(path):
    if os.path.isdir(path):
        try:
            shutil.rmtree(path)
        except Exception as e:
            raise e
        return True
    else:
        try:
            os.remove(path)
        except Exception as e:
            raise e
        return True
        

def main():
    profiles = get_profiles()
    choose = choose_profiles(profiles)
    installs = profiles_to_install(profiles, choose)

    for i, v in enumerate(installs):
        install(v)


if __name__ == '__main__':
    main()
