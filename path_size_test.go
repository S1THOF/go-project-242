package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize_file1(t *testing.T) {
	res, err := GetPathSize("testdata/file1.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "62289B	testdata/file1.csv", res)
}

func TestGetSize_file2(t *testing.T) {
	res, err := GetPathSize("testdata/file2.csv", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "62388B	testdata/file2.csv", res)
}

func TestGetSize_testdir(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", false, false, false)
	assert.NoError(t, err)
	assert.Equal(t, "124677B	testdata/testdir", res)
}

func TestGetSize_file1_human(t *testing.T) {
	res, err := GetPathSize("testdata/file1.csv", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, "60.8KB	testdata/file1.csv", res)
}

func TestGetSize_file2_human(t *testing.T) {
	res, err := GetPathSize("testdata/file2.csv", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, "60.9KB	testdata/file2.csv", res)
}

func TestGetSize_testdir_human(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", false, true, false)
	assert.NoError(t, err)
	assert.Equal(t, "121.8KB	testdata/testdir", res)
}

func TestGetSize_hiddendir_all_false(t *testing.T) {
	res, err := GetPathSize("testdata/.hiddendir", false, false, false)
	assert.NoError(t, err)
	// Все файлы внутри скрытые → игнорируются при all=false
	assert.Equal(t, "0B\ttestdata/.hiddendir", res)
}

func TestGetSize_hiddendir_all_true(t *testing.T) {
	res, err := GetPathSize("testdata/.hiddendir", false, false, true)
	assert.NoError(t, err)
	// Один скрытый файл размером 62388B
	assert.Equal(t, "62388B\ttestdata/.hiddendir", res)
}

func TestGetSize_testdir_recursive_true(t *testing.T) {
	res, err := GetPathSize("testdata/testdir", true, true, true)
	assert.NoError(t, err)
	assert.Equal(t, "243.5KB\ttestdata/testdir", res)
}

func TestFormat_b(t *testing.T) {
	assert.Equal(t, "1000B", FormatSize(1000))
}

func TestFormat_KB(t *testing.T) {
	assert.Equal(t, "1.0KB", FormatSize(1024))
}

func TestFormat_MB(t *testing.T) {
	assert.Equal(t, "1.0MB", FormatSize(1048576))
}
