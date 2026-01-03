package utils

import "strings"

// ToLower converte string para minúsculas (compatibilidade)
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Contains verifica se uma string contém outra (case-insensitive)
func Contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
