#! /usr/bin/env bash
set -eu

touch   '/opt/tarantool/storage.log'
tail -F '/opt/tarantool/storage.log' &
exec tarantool '/usr/local/bin/tarantool-entrypoint.lua' '/opt/tarantool/app.lua' >> '/dev/null'
