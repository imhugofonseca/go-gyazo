#!/bin/bash
echo "SetEnv GYAZO_HOST ${GYAZO_HOST}" >> /etc/apache2/conf.d/gyazo.conf
/usr/sbin/apache2ctl -D FOREGROUND