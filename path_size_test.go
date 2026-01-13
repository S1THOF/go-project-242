package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize_file1(t *testing.T) {
	res, err := GetPathSize("testdata/file1.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "62289B	testdata/file1.csv")
}

func TestGetSize_file2(t *testing.T) {
	res, err := GetPathSize("testdata/file2.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "62388B	testdata/file2.csv")
}

func TestGetSize_testdir(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "124677B	testdata/testdir")
}

func TestGetSize_file1_human(t *testing.T) {
	res, err := GetPathSize("testdata/file1.csv", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "60.8KB	testdata/file1.csv")
}

func TestGetSize_file2_human(t *testing.T) {
	res, err := GetPathSize("testdata/file2.csv", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "60.9KB	testdata/file2.csv")
}

func TestGetSize_testdir_human(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, res, "121.8KB	testdata/testdir")
}

func TestFormat_b(t *testing.T) {
	res := FormatSize(1000.0)
	assert.Equal(t, res, "1000.0B")
}

func TestFormat_KB(t *testing.T) {
	res := FormatSize(1024)
	assert.Equal(t, res, "1.0KB")
}

func TestFormat_MB(t *testing.T) {
	res := FormatSize(1048576)
	assert.Equal(t, res, "1.0MB")
}
