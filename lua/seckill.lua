local productId = ARGV[1]
local userId = ARGV[2]
local stockKey = KEYS[1]..productId
local orderKey = KEYS[2]..productId
local res = redis.call('get', stockKey)
if (tonumber(res, 10) <= 0) then
    return 1
end
if redis.call("sismember", orderKey, userId) == 1 then
    return 2
end
redis.call('incrby', stockKey, -1)
redis.call('sadd', orderKey, userId)
return 0