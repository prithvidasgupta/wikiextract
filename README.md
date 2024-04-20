# wikiextract
A tool to get the plainest text out of Wikipedia XML dumps. The output is a CSV file. Inspired by [attardi/Wikiextractor](https://github.com/attardi/wikiextractor)

Example usage:

```
wikiextract extract -f ./samplewiki.xml.bz2 -o out.csv
```

The above command is comparable to `wikiextractor ./sample.xml.bz2 --no-templates -ns ns0` of [attardi/Wikiextractor](https://github.com/attardi/wikiextractor)


### Benchmark device
Windows WSL -- 11th Gen Intel(R) Core(TM) i5-11300H @ 3.10GHz   3.11 GHz RAM 16.0 GB

### Benchmarking with Simple EN wiki dump

| wikiextract| wikiextractor|
|------------|--------------|
|real    0m45.275s|real    1m26.966s|
|user    1m37.802s|user    4m25.311s|
|sys     0m6.076s|sys     0m25.905s|

### Benchmarking with BN wiki dump

| wikiextract| wikiextractor|
|------------|--------------|
|real    1m36.287s |real    2m45.072s|
|user    2m59.901s|user    6m0.504s|
|sys     0m7.948s|sys     0m35.705s|

**Current Goal**: To achieve parity with [attardi/Wikiextractor](https://github.com/attardi/wikiextractor)

### Why do this at all?

I am trying to learn Go and I saw there were no Wiki XML dump extractors in Go, so might as well do it.
