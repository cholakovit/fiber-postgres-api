package queries

import (
    "fmt"
)

type DBInterface interface {
    Connect() (string, error)
}

type MockDB struct{}

func (m MockDB) Connect() (string, error) {
    return "Connected to mock DB", nil
}

type Initializer struct {
    Connect func() (DBInterface, error)
}

func main() {
    mockDB := MockDB{}
    
    initializer := Initializer{
        Connect: func() (DBInterface, error) {
            return mockDB, nil
        },
    }
    
    db, err := initializer.Connect()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println(db.Connect())
}
