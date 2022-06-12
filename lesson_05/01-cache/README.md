# Cache

Implement a simple cache, that holds values with type `string` and allows to retrieve them using keys of time `stings`. Key/value pairs can expire if given a deadline using an appropriate method. 

The initial structure is given. You should add appropriate fields, implement the constructor function and methods with given signatures. 

`k, ok := Get(key string)` - returns the value associated with the key and the boolean `ok` (true if exists, false if not), if the deadline of the key/value pair has not been exceeded yet


`Put(key string, value string)` places a value with an associated key into cache. Value put with this method never expired(have infinite deadline). Putting into the existing key should overwrite the value

`PutTill(key string, value string, deadline time.Time)`
Should do the same as `Put`, but the key/value pair should expire by given deadline

`Keys() []string` should return the slice of existing (non-expired keys)

The expiration should, and probably cannot in our case, be implemented asynchronously, so think of a way to clean up the records when needed

NOTE: receiver is a placeholder, replace appropriately
