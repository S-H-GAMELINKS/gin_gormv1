File.open("user_insert.sql", "w") do |file|
  10000.times do |i|
    file.flush << "insert into users values(#{i}, 'S-#{i}');\n"
  end
end

File.open("tweet_insert.sql", "w") do |file|
  10000.times do |i|
    file.flush << "insert into tweets values(#{i}, 'S-#{i}', #{rand(1..10000)});\n"
  end
end
