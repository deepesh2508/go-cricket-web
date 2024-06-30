package errs

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/deepesh2508/go-cricket-web/env"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	errorMessages = make(map[string]*ErrorResponse)
	db            *sql.DB
)

type ErrorResponse struct {
	H    int
	C    string
	M    string
	Args []interface{}
}

func (e *ErrorResponse) GetErrorMessage(c *gin.Context) string {
	msg := e.M
	args := e.Args

	n := strings.Count(msg, "%v")
	m := len(args)

	if n == 0 {
		return msg
	}

	if n <= m {
		return fmt.Sprintf(msg, args[:n]...)
	}

	for i := m; i < n; i++ {
		args = append(args, "")
	}

	return fmt.Sprintf(msg, args[:n]...)
}

func NewError(httpcode int, code string, message string) *ErrorResponse {
	if _, ok := errorMessages[code]; ok {
		panic(fmt.Sprintf("duplicate error code: %v", code))
	}

	err := &ErrorResponse{
		H: httpcode,
		C: code,
		M: message,
	}

	errorMessages[code] = err
	return err
}

func ErrorsCacheRefresh() {
	cacheTTL, err := time.ParseDuration(env.ENV.CACHE_TTL)
	if err != nil {
		// if invalid duration value, default to 15 minutes
		cacheTTL = 15 * time.Minute
	}

	for {
		// sleep for some time before getting all the values again
		time.Sleep(cacheTTL)

		// get all the error values from db
		GetErrorsFromDB()
	}
}

func GetErrorsFromDB() {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", env.ENV.DATABASE_URL)
		if err != nil {
			panic("failed to connect to database: " + err.Error())
		}
	}

	for _, e := range errorMessages {
		var message string
		query := `
			SELECT message
			FROM error_messages
			WHERE code = $1 AND active = true
			LIMIT 1;
		`
		err := db.QueryRow(query, e.C).Scan(&message)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("no active error message found for code: %s\n", e.C)
			} else {
				fmt.Printf("error fetching message for code %s: %v\n", e.C, err)
			}
			continue
		}
		e.M = message
	}
}
