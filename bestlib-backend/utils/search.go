package utils

type BookInfo struct {
    Title  string
    Author string
    Year   string
    Status string
}

func ExtractBookInfo(searchTerm string) BookInfo {
    return BookInfo{
        Title:  searchTerm,
        Author: searchTerm,
        Year:   searchTerm,
        Status: searchTerm,
    }
}
