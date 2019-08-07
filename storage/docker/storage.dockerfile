from tarantool/tarantool:2.1

COPY './storage/storage.lua' '/opt/tarantool/app.lua'
CMD tarantool '/usr/local/bin/tarantool-entrypoint.lua' '/opt/tarantool/app.lua'
