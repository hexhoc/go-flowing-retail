FROM postgres:latest
# Set ru_RU local (default is en_US)
# RUN localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8
# ENV LANG ru_RU.utf8

COPY config/postgres/migration /workdir/migration