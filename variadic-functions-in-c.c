
int  sum (int count,...) {
    
  int sum =0;
  
  va_list n;
  va_start( n, count );
  
  for( int i = 0; i < count; i++ )
    sum += va_arg( n, int );
    
  va_end( n );

  return sum;
}

int min(int count,...) {
    
    int min;
    int initialized = 0;
    va_list n;
    va_start( n, count );
  
    for( int i = 0; i < count; i++ ) {
       int k;
       k = va_arg( n, int ); 
       if (initialized == 0 || min > k) {
           min = k;
           initialized = 1;
       }
    }   
    
    va_end( n );

    return min;
}

int max(int count,...) {
    int max;
    int initialized = 0;
    va_list n;
    va_start( n, count );
  
    for( int i = 0; i < count; i++ ) {
       int k;
       k = va_arg( n, int ); 
       if (initialized == 0 || max < k) {
           max = k;
           initialized = 1;
       }
    }   
    va_end( n );

    return max;
}

