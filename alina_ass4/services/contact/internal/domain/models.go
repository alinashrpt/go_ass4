package domain

type Contact struct {
    ID           int
    FullName     string
    PhoneNumber  string
}

type Group struct {
    ID   int
    Name string
    Contacts []*Contact 
}
