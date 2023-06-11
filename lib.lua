#!lua name=lib
local function zsequence(keys, args)
   local lower_bound, upper_bound = unpack(args)
   local values = {}

   for i, k in ipairs(keys) do
      local value = redis.call('ZRANGE', k, lower_bound, upper_bound, 'BYSCORE', 'LIMIT', 0, 1, 'WITHSCORES')

      if #value > 0 then
         value = value[2]
         lower_bound = value + 1
      else
         value = false
      end

      table.insert(values, value)
   end

   return values
end

local function zvalue(keys, args)
   local lower_bound, upper_bound = unpack(args)
   local value = redis.call('ZRANGE', keys[1], lower_bound, upper_bound, 'BYSCORE', 'LIMIT', 0, 1, 'WITHSCORES')

   return value[2]
end

redis.register_function('zvalue', zvalue)
redis.register_function('zsequence', zsequence)

