package main

import (
	"hash/crc32"
	"strconv"
	"strings"
	"sync"
)

const shortURLPrefix = "http://tinyurl.com/"

type Codec struct {
	m sync.Map
}

func Constructor535() Codec {
	return Codec{
		m: sync.Map{},
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	n := crc32.ChecksumIEEE([]byte(longUrl))
	if _, ok := c.m.Load(n); !ok {
		c.m.Store(n, longUrl)
	}
	return shortURLPrefix + strconv.FormatUint(uint64(n), 16)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	id := strings.TrimPrefix(shortUrl, shortURLPrefix)
	n, _ := strconv.ParseUint(id, 16, 32)
	v, ok := c.m.Load(uint32(n))
	if !ok {
		return ""
	}
	return v.(string)
}
