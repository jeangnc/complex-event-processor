#!lua name=mylib
local function load_sequence(keys, args)
   local lower_bound, upper_bound = unpack(args)
   local values = {}

   for i, k in ipairs(keys) do
      local value = {}

      if lower_bound then
         value = redis.call('ZRANGE', k, lower_bound, upper_bound, 'BYSCORE', 'LIMIT', 0, 1, 'WITHSCORES')

         if #value > 0 then
            lower_bound = value[2] + 1
         else
            lower_bound = false
         end
      end

      table.insert(values, value)
   end

   return values
end

redis.register_function('load_sequence', load_sequence)
