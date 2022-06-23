package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWorld(t *testing.T) {
// 	result := HelloWorld("Setiawan")
// 	if result != "Hello Teguh" {
// 		// panic("Bukan Hello Teguh")
// 		// t.Fail()
// 		t.Error("Result must be 'Hello Teguh' ")
// 	}

// 	fmt.Println("TestHelloWorld Teguh Done")
// }

// func TestHelloWorldTeguh(t *testing.T) {
// 	result := HelloWorld("Teguh")
// 	if result != "Hello Teguh" {
// 		// panic("Bukan Hello Teguh")
// 		// t.FailNow()
// 		t.Fatal("Resutl must be 'Hello Teguh'")
// 	}
// 	fmt.Println("TestHelloWorldTeguh Teguh Done")
// }

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Teguh")
	assert.Equal(t, "Hello Teguh", result, "Result must be 'Hello Teguh'")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequireTeguh(t *testing.T) {
	result := HelloWorld("Teguh")
	require.Equal(t, "Hello Teguh", result)
	fmt.Println("TestHelloWorldTeguh Teguh Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on linux")
	}

	result := HelloWorld("Setiawan")
	require.Equal(t, "Hello Teguh", result)
}

func TestSubTest(t *testing.T) {
	t.Run("Teguh", func(t *testing.T) {
		result := HelloWorld("Teguh")
		require.Equal(t, "Hello Teguh", result)
	})
	t.Run("Setiawan", func(t *testing.T) {
		result := HelloWorld("Setiawan")
		require.Equal(t, "Hello Setiawan", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Teguh",
			request:  "Teguh",
			expected: "Hello Teguh",
		}, {
			name:     "Setiawan",
			request:  "Setiawan",
			expected: "Hello Setiawan",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Teguh")
	}
}

func BenchmarkHelloWorldSetiawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Setiawan")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Teguh", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Teguh")
		}
	})
	b.Run("Setiawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Setiawan")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Teguh",
			request:  "Teguh",
			expected: "Hello Teguh",
		}, {
			name:     "Setiawan",
			request:  "Setiawan",
			expected: "Hello Setiawan",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		}, {
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}
