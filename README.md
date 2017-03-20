# TCACHE or PMRC

[![GoDoc](https://godoc.org/github.com/wind85/tcache?status.svg)](https://godoc.org/github.com/wind85/tcache)
[![Build Status](https://travis-ci.org/wind85/tcache.svg?branch=master)](https://travis-ci.org/wind85/tcache)
[![Coverage Status](https://coveralls.io/repos/github/wind85/tcache/badge.svg?branch=master)](https://coveralls.io/github/wind85/tcache?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
### Poor Man Redis Cache also known a timed cache or tcache
This is a small package the provides a timed, goroutine safe cache with auto expiration
for the elements. It provides the following features:

- New accept 2 parameters, the first one a time.Duration expressed in minutes, that dictates,
  every time the map is evaluated for values to be evicted, and a second value that sets the,
  expiration time of every element in the map.
- Put method accept any value and as key takes a string, if Put is called more than once
  with the same value, the same rules that apply to a standard map apply here as well.
- Get retrieves the given value by key and returns the value and a bool, if the key isn't
  there either because it is expired or there has not been one a false bool is returned and
  a nil. In case the value is found, an interface with a valid value is returned and bool with
  true as a value.

### How to use it

Pretty simple, there is only one method to create a new parser just call 
```
  cache := tcache.New(5,10) 
  // every 5 minutes elements are checked to be evicted, 10 minutes is the expiration time set on every value
```
To put and retrieve value to like so:
```
  cache.Put("keyname","value") // any value can be put inside the cache
  value, ok := cache.Get("keyname")
  if !ok {
      log.Println("value not expired")
  }
```
The value expires after 10 minutes ( in this example ) and it is automatically evicted.

#### Philosophy
This software is developed following the "mantra" keep it simple, stupid or better known as
KISS. Something so simple like a cache with auto eviction should not required over engineered 
solutions. Though it provides most of the functionality needed by generic configuration files, 
and most important of all meaning full error messages.

#### Disclaimer
This software in alpha quality, don't use it in a production environment, it's not even
completed.

#### Thank You Notes
I thank myself since I wrote this because I didn't want to use redis for such a simple case.
