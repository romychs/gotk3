#!/usr/bin/env bash

# if [ -z  "$1" ]; then
    PREFIX=/usr

    # FreeBSD
    if [[ $(echo "$OSTYPE" | tr "[:upper:]" "[:lower:]") == "freebsd"* ]]; then
        PREFIX="${PREFIX}/local"
    # Linux OS
    # elif [[ "$OSTYPE" == "linux-gnu" ]]; then
    # Mac OSX
    # elif [[ "$OSTYPE" == "darwin"* ]]; then
    # POSIX compatibility layer and Linux environment emulation for Windows
    # elif [[ "$OSTYPE" == "cygwin" ]]; then
    # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
    # elif [[ "$OSTYPE" == "msys" ]]; then
    # Windows
    # elif [[ "$OSTYPE" == "win32" ]]; then
    # else
            # Unknown.
    fi
# else
#    export PREFIX=$1
# fi

if [ "$(id -u)" != "0" ]; then
    # Make sure only root can run our script
    echo "This script must be run as root" 1>&2
    exit 1
fi

SCHEMA_PATH=${PREFIX}/share/glib-2.0/schemas
echo "Uninstalling gsettings schema from ${SCHEMA_PATH}"

rm ${SCHEMA_PATH}/org.d2r2.gotk3.cool_app_1.gschema.xml
# Redirect output to /dev/null help on some linux distributions (redhat), which produce
# lot of warnings about "Schema ... are depricated." not related to application.
glib-compile-schemas ${SCHEMA_PATH}/ 2>/dev/null
