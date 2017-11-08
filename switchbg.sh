export IMG=$HOME/go/src/github.com/jamiebarnett/gobackground/img/

sqlite3 ~/Library/Application\ Support/Dock/desktoppicture.db "update data set value = '$IMG/octopus.jpg'" && killall Dock 