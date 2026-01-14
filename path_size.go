package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func calcSize(path string, recursive, all bool) (int64, error) {
	fileinfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !fileinfo.IsDir() {
		if fileinfo.Mode().IsRegular() {
			return fileinfo.Size(), nil
		}
		return 0, nil
	}

	// Это директория
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range entries {
		if !shouldInclude(entry.Name(), all) {
			continue
		}

		if entry.IsDir() {
			if recursive {
				fullPath := filepath.Join(path, entry.Name())
				size, err := calcSize(fullPath, recursive, all)
				if err != nil {
					return 0, err
				}
				total += size
			}
		} else {
			// Обычный файл
			info, err := entry.Info()
			if err != nil {
				return 0, err
			}
			if info.Mode().IsRegular() {
				total += info.Size()
			}
		}
	}
	return total, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := calcSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	var result string
	if human {
		result = fmt.Sprintf("%s\t%s", FormatSize(float64(size)), path)
	} else {
		result = fmt.Sprintf("%vB\t%s", size, path)
	}
	return result, nil
}

func FormatSize(size float64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0

	for size >= 1024 && unitIndex < len(units)-1 {
		size /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%.0fB", size)
	}
	return fmt.Sprintf("%.1f%s", size, units[unitIndex])
}

func shouldInclude(name string, all bool) bool {
	return all || !strings.HasPrefix(name, ".")
}
