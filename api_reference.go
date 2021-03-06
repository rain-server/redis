package red

type APIReference struct {
	Strings struct{
		GET GET
		SET SET
		SETNX SETNX
		SETXX SETXX
		DEL DEL
		DECR DECR
		INCR INCR
	}
	Lists struct{
		LPUSH LPUSH
		LPUSHX LPUSHX
		RPUSH RPUSH
		RPUSHX RPUSHX
		LPOP LPOP
		LPOPCount LPOPCount
		RPOP RPOP
		RPOPCount RPOPCount
		LRANGE LRANGE
		BRPOPLPUSH BRPOPLPUSH
	}
	Hashes struct{
		HGET HGET
		HMGET HMGET
		HSET HSET
	}
	Stream struct{
		XADD XADD
		XLEN XLEN
		XRANGE XRANGE
		XDEL XDEL
		XPEDING XPEDING
		XREAD XREAD
		XREADGROUP XREADGROUP
		XACK XACK
		XGROUPCreate XGROUPCreate
		XTRIM XTRIM
	}
}
