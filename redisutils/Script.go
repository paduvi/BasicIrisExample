package redisutils

const ViewItemByUserIdScript string =`
	local exist = redis.call("ZADD", KEYS[1], ARGV[3], ARGV[2])
	if exist == 1 then
		redis.call("ZINCRBY", "user:list", 1, ARGV[1])
	end
	redis.call("EXPIRE", KEYS[1], 120)
	redis.call("SETEX", KEYS[2], 360, ARGV[3])
	`
