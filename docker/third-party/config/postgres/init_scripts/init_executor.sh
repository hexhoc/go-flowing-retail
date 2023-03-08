#!/bin/sh
export PGUSER=postgres

# ============================================================================
# INITIAL
# ============================================================================

psql -v ON_ERROR_STOP=1 -f '/workdir/migration/1_init.sql' || exit $?
export PGDATABASE=product_service
psql -v ON_ERROR_STOP=1 -f '/workdir/migration/2_product_service.sql' || exit $?
export PGDATABASE=postgres