register_function:
	cat lib.lua | docker exec -i redis redis-cli -x FUNCTION LOAD REPLACE
