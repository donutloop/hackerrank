SELECT DISTINCT city FROM station WHERE city REGEXP '^([aeiou]{1,1}).*([aeiou]{1,1})$' = 0;