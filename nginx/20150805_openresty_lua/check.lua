local cjson = require("cjson")
local funcs = require("funcs")

local method = ngx.req.get_method()

if method == "GET" then
    local e = ngx.var.arg_e
    if not e or (tonumber(e) ~= 0 and tonumber(e) < os.time()) then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="e invalid"}))
        return
    end

    local ret = funcs.check_token()
    if ret < 0 then
        return
    end
elseif method == "POST" or method == "DELETE" then
    local we = ngx.var.arg_we
    if not we or tonumber(we) < os.time() then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="we invalid"}))
        return
    end
    local ret = funcs.check_token()
    if ret < 0 then
        return
    end
else
    ngx.status = ngx.HTTP_BAD_REQUEST
    ngx.say(cjson.encode({error="method invalid"}))
    return
end

local fid = string.gsub(ngx.var.uri,"/f/([0-9a-z,]*).*","%1")
local query = string.gsub(ngx.var.uri,"/f/(.*)","%1").."?"..ngx.var.query_string
local master = "/masters/"..fid

local res = ngx.location.capture(master)
if res.status ~= ngx.HTTP_MOVED_PERMANENTLY then
    ngx.status = res.status
    ngx.say(res.body)
    return
end

if res.status ~= ngx.HTTP_MOVED_PERMANENTLY then
    ngx.status = res.status
    ngx.say(res.body)
    return
end

ngx.var.volume = res.header["Location"]