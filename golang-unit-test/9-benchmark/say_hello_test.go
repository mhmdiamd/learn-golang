package benchmark

import "testing"

func BenchmarkHelloWorld(b *testing.B) {
	for i := 1; i < b.N; i++ {
		SayHello("Ilham")
	}
}

// Table Benchmark

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Muhamad)",
			request: "Muhamad",
		},
		{
			name:    "HelloWorld(Ilham)",
			request: "Ilham",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}
