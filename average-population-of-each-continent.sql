select COUNTRY.Continent, floor(avg(CITY.POPULATION)) from CITY join COUNTRY on COUNTRY.Code = CITY.CountryCode group by COUNTRY.Continent ;