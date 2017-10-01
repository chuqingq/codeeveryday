# logrotate
if [ ! -f ${SHELL_DIR}/logrotate ]; then
    LOG_FILE=$(dirname ${SHELL_DIR})/log/channel.0
    sed "s|LOG_FILE|${LOG_FILE}|g" ${SHELL_DIR}/../../logrotate.tpl > ${SHELL_DIR}/logrotate
    crontab << _END
`crontab -l|grep -v ${SHELL_DIR}/logrotate`
*/10 * * * * /usr/sbin/logrotate ${SHELL_DIR}/logrotate
`echo`
_END
    echo "generate logrotate and add crontab success"
fi


