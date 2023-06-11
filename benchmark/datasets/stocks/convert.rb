require 'csv'
require 'date'

CSV.open("ALL_210019836.csv", "r") do |csv1|
  CSV.open("ALL_210019836_converted.csv", 'w') do |csv2|
    csv1.each do |r|
      r << DateTime.parse("1970-01-01T#{r[4]}+00:00").to_time.to_i
      csv2 << r
    end
  end
end
