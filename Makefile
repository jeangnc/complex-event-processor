register_function:
	cat lib.lua | redis-cli -x FUNCTION LOAD REPLACE
