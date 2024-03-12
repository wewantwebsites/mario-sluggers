package models

type Character struct { 
    ID  int `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Ability     string  `json:"ability"`
    Team        string  `json:"team"`
    Stats       Stats   `json:"stats"`
}

type Stats struct { 
    Pitch   int `json:"pitch"`
    Bat     int `json:"bat"`
    Field   int `json:"field"`
    Run     int `json:"run"`
}
