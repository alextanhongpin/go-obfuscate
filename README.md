# go-obfuscate

How can we encrypt the int primary key of a relational database (e.g. MySQL)?

1. aes-siv (the fastest choice if you have to do it, used by connect2id to encrypt user id [^1]). Also mentioned in [^5].
2. aes-gcm (not an option, because the salt is random)
3. hashid (possible, but slow)
4. None of the above, just generate a uuid and store in the db as the secondary key [^4].


TL;DR; stick with option 4. How others do it [^2] and [^3].


The issue with all the other approaches is that the logic to generate such identifier is tied to the application. Consider the scenario where you need to generate a unique user identifier to be send to analytics tool. Now it becomes a problem to identify which id belongs to which user if they are generated through the app.


Instead, if they are stored in the database, queries can still be done directly to access them. Note that if you don't have to return the ids to the end-users, then this article can be disregarded.

For some other scenarios like pagination token etc, using AES-SIV works better over AES-GCM too because the result will always be consistent when encrypting with the same payload, which could be useful for caching etc.


## Benchmark

```bash
➜  go-obfuscate go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/alextanhongpin/go-obfuscate
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkHashID-4         131371              8760 ns/op
BenchmarkDaead-4          283908              3939 ns/op
BenchmarkAESGCM-4         364515              2951 ns/op
PASS
ok      github.com/alextanhongpin/go-obfuscate  9.290s
```


[^1]: https://connect2id.com/blog/deterministic-encryption-with-aes-siv
[^2]: https://andrew.carterlunn.co.uk/programming/2020/05/17/encrypting-integer-primary-keys.html
[^3]: https://medium.com/@patrickfav/a-better-way-to-protect-your-database-ids-a33fa9867552
[^4]: https://stackoverflow.com/questions/32795998/hiding-true-database-object-id-in-urls
[^5]: https://www.reddit.com/r/cryptography/comments/bit7k8/unsatisfied_with_hashids_ive_created_sound/
