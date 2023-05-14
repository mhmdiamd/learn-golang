package assertion

// Assert berfungsi jika ada error testing maka program akan tetap di jalankan
// Require berfungsi jika ada error testing maka program akan diberhentikan dan kode dibawahnya tidak dijalankan

import (
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {

	if runtime.GOOS == "darwin" {
		t.Skip("Can not run in mac OS")
	}

}
