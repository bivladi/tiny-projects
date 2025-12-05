package genericcache_test

import (
	cache "learngo/genericcache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const defaultTtl = time.Minute * 5
const defaultMaxSize = 5

func TestCache(t *testing.T) {
	c := cache.New[int, string](defaultTtl, defaultMaxSize)

	_ = c.Upsert(1, "one")
	v, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, "one", v)

	_ = c.Upsert(2, "two")
	v, found = c.Read(2)
	assert.True(t, found)
	assert.Equal(t, "two", v)

	c.Delete(1)
	v, found = c.Read(1)
	assert.False(t, found)

	v, found = c.Read(2)
	assert.True(t, found)
	assert.Equal(t, "two", v)

	_ = c.Upsert(2, "two two")
	v, found = c.Read(2)
	assert.True(t, found)
	assert.Equal(t, "two two", v)
}

func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string](defaultTtl, defaultMaxSize)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		_ = c.Upsert(6, "six")
	})
	t.Run("write value", func(t *testing.T) {
		t.Parallel()
		_ = c.Upsert(6, "value")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](defaultTtl, defaultMaxSize)
	_ = c.Upsert("key", "value")
	got, found := c.Read("key")
	assert.True(t, found)
	assert.Equal(t, "value", got)

	time.Sleep(10 * time.Millisecond)

	got, found = c.Read("key")
	assert.False(t, found)
	assert.Equal(t, "", got)
}

func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	// Give it a TTL long enough to survive this test
	c := cache.New[int, int](time.Minute, 3)

	_ = c.Upsert(1, 1)
	_ = c.Upsert(2, 2)
	_ = c.Upsert(3, 3)

	got, found := c.Read(1)
	assert.True(t, found)
	assert.Equal(t, 1, got)

	// Update 1, which will no longer make it the oldest
	_ = c.Upsert(1, 10)

	// Adding a fourth element will discard the oldest - 2 in this case.
	_ = c.Upsert(4, 4)

	// Trying to retrieve an element that should've been discarded by now.
	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)
}
