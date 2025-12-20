local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window_micro = tonumber(ARGV[2]) 
local current_time = tonumber(ARGV[3])
local member = ARGV[4]           
local window_seconds = tonumber(ARGV[5])

redis.call('ZREMRANGEBYSCORE', key, 0, current_time - window_micro)

local current_count = redis.call('ZCARD', key)

if current_count < limit then
    redis.call('ZADD', key, current_time, member)
    redis.call('EXPIRE', key, window_seconds)
    return {1, 0, current_count + 1} -- Allowed
else
    local oldest = redis.call('ZRANGE', key, 0, 0, 'WITHSCORES')
    local retry_after_micro = 0
    
    if #oldest > 0 then
        retry_after_micro = tonumber(oldest[2]) + window_micro - current_time
    end
    
    if retry_after_micro < 0 then retry_after_micro = 0 end
    return {0, retry_after_micro, current_count} -- Denied
end