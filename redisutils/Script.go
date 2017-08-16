package redisutils

const ViewItemByUserIdScript string =`
	local exist = redis.call("ZADD", KEYS[1], ARGV[3], ARGV[2])
	if exist == 1 then
		redis.call("ZINCRBY", "user:list", 1, ARGV[1])
	end
	redis.call("EXPIRE", KEYS[1], 259200)
	redis.call("SETEX", KEYS[2], 259200, ARGV[3])
	`

const RemoveOldHistoryScript string = `
	redis.call("ZREM", KEYS[1], ARGV[2])
	local len = redis.call("ZCARD", KEYS[1])
	if len == 0 then
		redis.call("DEL", KEYS[1])
		redis.call("ZREM", "user:list", ARGV[1])
	else
		redis.call("ZINCRBY", "user:list", -1, ARGV[1])
	end
`