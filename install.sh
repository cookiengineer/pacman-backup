#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" > /dev/null && pwd)";


if [[ "$(which node 2>/dev/null)" == "" ]]; then
	sudo pacman -S --needed --noconfirm nodejs;
fi;


if [ ! -x /usr/bin/pacman-backup ]; then

	sudo cp "$DIR/pacman-backup.js" /usr/bin/pacman-backup;
	sudo chmod +x /usr/bin/pacman-backup;

fi;

