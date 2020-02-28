package transaction

// This file contains utility code for testing commands.

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/pingcap-incubator/tinykv/kv/tikv"
	"github.com/pingcap-incubator/tinykv/kv/tikv/inner_server"
	"github.com/pingcap-incubator/tinykv/kv/tikv/transaction/mvcc"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/stretchr/testify/assert"
)

// testBuilder is a helper type for running command tests.
type testBuilder struct {
	t      *testing.T
	server *tikv.Server
	// mem will always be the backing store for server.
	mem *inner_server.MemInnerServer
	// Keep track of timestamps.
	prevTs uint64
}

// kv is a type which identifies a key/value pair to testBuilder.
type kv struct {
	cf string
	// The user key (unencoded, no time stamp).
	key []byte
	// Can be elided. The builder's prevTS will be used if the ts is needed.
	ts uint64
	// Can be elided in assertion functions. If elided then testBuilder checks that the value has not changed.
	value []byte
}

func newBuilder(t *testing.T) testBuilder {
	mem := inner_server.NewMemInnerServer()
	server := tikv.NewServer(mem)
	return testBuilder{t, server, mem, 99}
}

// init sets values in the test's DB.
func (builder *testBuilder) init(values []kv) {
	for _, kv := range values {
		ts := kv.ts
		if ts == 0 {
			ts = builder.prevTs
		}
		switch kv.cf {
		case engine_util.CfDefault:
			builder.mem.Set(kv.cf, mvcc.EncodeKey(kv.key, ts), kv.value)
		case engine_util.CfWrite:
			builder.mem.Set(kv.cf, mvcc.EncodeKey(kv.key, ts), kv.value)
		case engine_util.CfLock:
			builder.mem.Set(kv.cf, kv.key, kv.value)
		}
	}
}

func (builder *testBuilder) runRequests(reqs ...interface{}) []interface{} {
	var result []interface{}
	for _, req := range reqs {
		reqName := fmt.Sprintf("%v", reflect.TypeOf(req))
		reqName = strings.TrimPrefix(strings.TrimSuffix(reqName, "Request"), "*kvrpcpb.")
		fnName := "Kv" + reqName
		serverVal := reflect.ValueOf(builder.server)
		fn := serverVal.MethodByName(fnName)
		ctxtVal := reflect.ValueOf(context.Background())
		reqVal := reflect.ValueOf(req)

		results := fn.Call([]reflect.Value{ctxtVal, reqVal})

		assert.Nil(builder.t, results[1].Interface())
		result = append(result, results[0].Interface())
	}
	return result
}

// runOneCmd is like runCommands but only runs a single command.
func (builder *testBuilder) runOneRequest(req interface{}) interface{} {
	return builder.runRequests(req)[0]
}

func (builder *testBuilder) nextTs() uint64 {
	builder.prevTs++
	return builder.prevTs
}

// ts returns the most recent timestamp used by testBuilder as a byte.
func (builder *testBuilder) ts() byte {
	return byte(builder.prevTs)
}

// assert that a key/value pair exists and has the given value, or if there is no value that it is unchanged.
func (builder *testBuilder) assert(kvs []kv) {
	for _, kv := range kvs {
		var key []byte
		ts := kv.ts
		if ts == 0 {
			ts = builder.prevTs
		}
		switch kv.cf {
		case engine_util.CfDefault:
			key = mvcc.EncodeKey(kv.key, ts)
		case engine_util.CfWrite:
			key = mvcc.EncodeKey(kv.key, ts)
		case engine_util.CfLock:
			key = kv.key
		}
		if kv.value == nil {
			assert.False(builder.t, builder.mem.HasChanged(kv.cf, key))
		} else {
			assert.Equal(builder.t, kv.value, builder.mem.Get(kv.cf, key))
		}
	}
}

// assertLen asserts the size of one of the column families.
func (builder *testBuilder) assertLen(cf string, size int) {
	assert.Equal(builder.t, size, builder.mem.Len(cf))
}

// assertLens asserts the size of each column family.
func (builder *testBuilder) assertLens(def int, lock int, write int) {
	builder.assertLen(engine_util.CfDefault, def)
	builder.assertLen(engine_util.CfLock, lock)
	builder.assertLen(engine_util.CfWrite, write)
}
