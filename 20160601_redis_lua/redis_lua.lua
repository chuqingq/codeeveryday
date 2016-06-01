local res = redis.call('lrange', '1', 0, -1)

local newV = ''
for k,v in pairs(res) do
	newV = newV .. v .. ','
	redis.call('set', '1', newV)
end
return 'success'
-- redis-cli --eval redis_lua.lua