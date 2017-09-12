LOG_FILE {
    rotate 50
    size 200M
    dateext
    dateformat -%Y%m%d
    copytruncate
    notifempty
    missingok
    lastaction
        mv LOG_FILE-`date +%Y%m%d`{,-`date +%s`}
    endscript
}

