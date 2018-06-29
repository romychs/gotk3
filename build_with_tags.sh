#!/bin/sh

# Autodetected distributive specific versions
GTK=$(pkg-config --modversion gtk+-3.0 | tr . _| cut -d '_' -f 1-2)
GLib=$(pkg-config --modversion glib-2.0 | tr . _| cut -d '_' -f 1-2)

function usage()
{

    PROG=$(basename $0)
    SHORT_HELP="Usage: ${PROG} [options]
    Options:
      -GLib <version>   distributive specific GLib version.
      -GTK <version>    distributive specific GTK version.
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
            ENVIRONMENT=$VALUE
            ;;
        --GLib)
            DB_PATH=$VALUE
            ;;
        *)
            echo "ERROR: unknown parameter \"$PARAM\""
            usage
            exit 1
            ;;
    esac
    shift
done

echo "GOTK3 build across GLib ${GLib}, GTK ${GTK} ..."
go build -v -tags "glib_${GLib} gtk_${GTK}" ./...


