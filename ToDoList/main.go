package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
  // Buat UUID baru
	id := uuid.New()

	// Ubah UUID menjadi string
	idStr := id.String()

	// Hapus tanda hubung dari UUID
	idStrNoHyphens := strings.ReplaceAll(idStr, "-", "")

	
	fmt.Println("UUID tanpa tanda hubung:", idStrNoHyphens)
}
