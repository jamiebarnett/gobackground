export IMG=$HOME/go/src/github.com/jamiebarnett/gobackground/img/
echo $IMG

sqlite3 ~/Library/Application\ Support/Dock/desktoppicture.db "update data set value = '$IMG/beautiful.jpg'" && killall Dock