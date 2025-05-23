package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func ReaderXmlGz(path string) (*os.File, *gzip.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("os.Open - %w", err)
	}

	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, nil, fmt.Errorf("gzip.NewReader - %w", err)
	}

	return file, reader, nil
}

func InArray[T comparable](list []T, include T) bool {
	for _, v := range list {
		if v == include {
			return true
		}
	}

	return false
}

func UniqueArray[T comparable](slice []T) []T {
	uniqMap := make(map[T]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]T, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}

	return uniqSlice
}

func VersionBytes(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}
