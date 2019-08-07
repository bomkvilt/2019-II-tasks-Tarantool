#!/usr/bin/env tarantool
local log  = require('log')
local json = require('json')

-- -----------| database |----------- --
box.cfg{
    listen=7011,
    log_format='plain',
    log='storage.log',
}

box.once('storage-1.0', function()
    log.info('storage bootstrapping')
    local storage = box.schema.create_space('storage')
    storage:create_index('primary', {
        type = 'HASH', parts = {1, 'string'}
    })
end)

-- -----------| helpers |----------- --

function getMessage(req, fields)
    local rawData = req:read()
    local ok, msg = pcall(json.decode, rawData)
    if not ok then
        return nil, rawData
    end

    for k, v in pairs(fields) do 
        if msg[v] == nil then
            return nil, rawData
        end
    end
    
    return msg, rawData
end

-- -----------| handlers |----------- --
function AddKey(req)
    local msg, raw = getMessage(req, {'key', 'value'})
    if msg == nil then
        log.info('add_key: incorrect message:. '..raw)
        return { status = 400 }
    end

    local ok, res = pcall(function() 
        return box.space.storage:insert{msg.key, msg.value}
    end)
    if not ok or res == ER_TUPLE_FOUND then
        log.info('add_key: key already exists')
        return { status = 409 }
    end
    
    log.info('add_key: success')
end

function UpdateKey(req)
    local key = req:stash('id')
    local msg, raw = getMessage(req, {'value'})
    if msg == nil then
        log.info('update_key: incorrect message ('..key..'):. '..raw)
        return { status = 400 }
    end

    local res = box.space.storage:update(key, {{'=', 2, msg.value}})
    if res == nil then
        log.info('get_key: key "'..key..'" not extsts')
        return { status = 404 }
    end
    log.info('update_key: success')
end

function GetKey(req)
    local key = req:stash('id')
    local res = box.space.storage:select(key)
    if #res ~= 1 then
        log.info('get_key: key "'..key..'" not extsts')
        return { status = 404 }
    end

    log.info('get_key: success')
    return { status = 200, body = json.encode(res[1][2]) }
end

function DeleteKey(req)
    local key = req:stash('id')
    if box.space.storage:delete(key) == nil then
        log.info('delete_key: key "'..key..'" not extsts')
        return { status = 404 }
    end

    log.info('delete_key: success')
    return { status = 200 }
end

-- -----------| server |----------- --
require('http.server').new('127.0.0.1', 7001)
    :route({ path = '/kv'    , method = 'POST'}, AddKey)
    :route({ path = '/kv/:id', method = 'GET' }, GetKey)
    :route({ path = '/kv/:id', method = 'PUT' }, UpdateKey)
    :route({ path = '/kv/:id', method = 'DELETE'}, DeleteKey)
    :start()
