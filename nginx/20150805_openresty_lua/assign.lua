local funcs = require("funcs")

local cjson = require("cjson")

if not ngx.var.arg_collection then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.say(cjson.encode({error="collection invalid"}))
    return
end
local ae = ngx.var.arg_ae
if not ae or tonumber(ae) < os.time() then
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.say(cjson.encode({error="ae invalid"}))
    return
end

local ret = funcs.check_token()
if ret < 0 then
    return
end

local res = ngx.location.capture("/masters"..ngx.var.request_uri)
ngx.status = res.status
if res.status ~= ngx.HTTP_OK then
    ngx.say(res.body)
    return
end
local res2 = cjson.decode(res.body)
res2.publicUrl = "http://192.168.13.165/f/"..res2.fid
res2.fid = nil 
res2.url = nil 
ngx.say(cjson.encode(res2))