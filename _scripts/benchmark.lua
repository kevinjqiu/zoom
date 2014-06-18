request = function()
    ip = math.random(256) .. "." .. math.random(256) .. "." .. math.random(256) .. "." .. math.random(256)
    return wrk.format("GET", "/geo/" .. ip)
end
