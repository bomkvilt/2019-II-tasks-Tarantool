from tarantool/tarantool:2.1

EXPOSE 80
COPY './storage/storage.lua' '/opt/tarantool/app.lua'
COPY './storage/launch.sh'   '/opt/tarantool/launch.sh'
CMD sh '/opt/tarantool/launch.sh'
