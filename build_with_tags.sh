#!/usr/bin/env bash

# Autodetected distributive specific versions
GLIB=$(pkg-config --modversion glib-2.0 | tr . _| cut -d '_' -f 1-2)
GTK=$(pkg-config --modversion gtk+-3.0 | tr . _| cut -d '_' -f 1-2)
GLIB_AUTO='(autodetected)'
GTK_AUTO=$GLIB_AUTO

function usage()
{

    PROG=$(basename $0)
    SHORT_HELP="Usage: ${PROG} [options]
    Options:
      --GLIB=<version>   distributive specific GLib version.
      --GTK=<version>    distributive specific GTK version.
      -h                Show this help message."

    echo "$SHORT_HELP"
    echo ""
}

while [ "$1" != "" ]; do
    PARAM=`echo $1 | awk -F= '{print $1}'`
    VALUE=`echo $1 | awk -F= '{print $2}'`
    case $PARAM in
        -h | /? | --help)
            usage
            exit
            ;;
        --GTK)
            GTK=$VALUE
            GTK_AUTO=''
            ;;
        --GLIB)
            GLIB=$VALUE
            GLIB_AUTO=''
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            usage
            exit 1
            ;;
    esac
    shift
done

echo "GOTK3 is building across GLIB ${GLIB}${GLIB_AUTO}, GTK ${GTK}${GTK_AUTO} ..."
go build -v -tags "glib_${GLIB} gtk_${GTK}" ./...


