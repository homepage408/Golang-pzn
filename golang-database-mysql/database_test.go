package golang_database_mysql

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	fmt.Println("Hai")
}

// export GO111MODULE=off
