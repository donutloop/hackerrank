import calendar

if __name__ == '__main__':
    n = raw_input().strip().split()
    
    
day = calendar.weekday(int(n[2]), int(n[0]), int(n[1]))    
    
if day == calendar.MONDAY:
    print "MONDAY"
elif day == calendar.TUESDAY:
    print "TUESDAY"
elif day == calendar.WEDNESDAY:
    print "WEDNESDAY"
elif day == calendar.THURSDAY:
    print "THURSDAY"
elif day == calendar.FRIDAY:
    print "FRIDAY"
elif day == calendar.SATURDAY:
    print "SATURDAY"
elif day == calendar.SUNDAY:
    print "SUNDAY"                
