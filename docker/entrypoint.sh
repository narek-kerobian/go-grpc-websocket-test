#!/bin/bash
set -e

# Load env params into system wide config
echo "Loading env parameters into system wide configuration"
env >> /etc/environment

# Start caddy and php-fpm
echo "Start supervisor daemon"
/usr/bin/supervisord -c /etc/supervisor/supervisord.conf

