local cjson = require("cjson")

local _M = {}

function _M.check_token()
    local crypt_uri = string.gsub(ngx.var.request_uri,"(.*)&t=.*", "%1")
    if not ngx.var.arg_t then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="t invalid: no t"}))
        return -1
    end
    local ak = string.gsub(ngx.var.arg_t, "(.*):.*", "%1")
    if not ak then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="t invalid: no ak"}))
        return -1
    end
    local token = string.gsub(ngx.var.arg_t, ".*:(.*)", "%1")
    if not token then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="t invalid: no token"}))
        return -1
    end
    local res = ngx.capture("/getsk", {args={ak=ak}})
    if res.status ~= ngx.HTTP_OK then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="t invalid: t check error"}))
        return -1
    end
    -- TODO 根据ak获取sk，增加location /getskbyak?ak=xxxx，校验IP只能是本地
    local expect_t = ngx.hmac_sha1("5GIaNcWym9MLGLdM",crypt_uri)
    local expect_t_hex = ndk.set_var.set_encode_hex(expect_t)
    if expect_t_hex ~= token then
        ngx.status = ngx.HTTP_BAD_REQUEST
        ngx.say(cjson.encode({error="t invalid: t check error"}))
        return -1
    end

    return 0
end

return _M
