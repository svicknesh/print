# Debug print structure to command line

## Usage

```go
type User struct {
    ID int `json:"id"`
    Username string `json:"username"`
}

user := User{ID: 1, Username: "Vanguard"}

print.JSON(user)

```
