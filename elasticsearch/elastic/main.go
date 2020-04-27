package main


var (
	subject Subject
	//indexName = "subject"
	//servers = []string{"http://localhost:9200/"}
)

//type Subject struct {
//	ID     int      `json:"id"`
//	Title  string   `json:"title"`
//	Genres []string `json:"genres"`
//}

const mapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`
