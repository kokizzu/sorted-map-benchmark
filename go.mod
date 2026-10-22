module modernc.org/benchmark-ordered-map

go 1.26.10

require (
	github.com/3xian/elephantlist v0.0.0-20140607022226-3d608579b1a8
	github.com/Michal-Leszczynski/btree v0.0.0-00010101000000-000000000000
	github.com/ajwerner/btree v0.0.0-20211221152037-f427b3e689c0
	github.com/biogo/store v0.0.0-20201120204734-aad293a2328f
	github.com/glenn-brown/skiplist v0.0.0-20121122060428-b6428f513193
	github.com/google/btree v1.1.2
	github.com/kellydunn/go-art v0.0.1
	github.com/petar/GoLLRB v0.0.0-20210522233825-ae3b015fd3e9
	github.com/ryszard/goskiplist v0.0.0-20150312221310-2dfbae5fcf46
	modernc.org/b v1.0.4
	modernc.org/b/v2 v2.1.0
)

require github.com/glenn-brown/ordinal v0.0.0-20121112042833-18b2aec8a941 // indirect

replace github.com/Michal-Leszczynski/btree => ./michal_leszczynski/btree/
