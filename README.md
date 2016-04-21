# dburl
Helper library for handling URI database connection strings for Go

## Usage

```go
import "github.com/evilmarty/dburl"

func main() {
  db, err := dburl.Open("mysql://user:pass@localhost/foobar")
  if err != nil {
    panic(err)
  }

  rows, err := db.Query("SELECT * FROM table")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  ...
}
```

## Todo

[ ] Verify parameters for MySQL
[ ] Add support for sqlite

