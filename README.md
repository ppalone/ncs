# NCS

Unofficial NCS (Nocopyrightsounds) Go client

## Installation

```bash
go get github.com/ppalone/ncs
```

## Usage

Search song

```go
// create a new client
c := ncs.NewClient(nil)

// search songs by "Tobu"
// query string can be any Song / Artist
q := "Tobu"
res, err := c.Search(context.Background(), q)
if err != nil {
  panic(err)
}

// get all songs res.Songs
fmt.Println(res.Songs)

// by default song results are paginated
// you can know if further results are available or not by res.HasNext bool
if (res.HasNext) {
  nextRes, err := res.Next(context.Background())
  if err != nil {
    panic(err)
  }

  // access further song search results
  fmt.Println(nextRes.Songs)
}
```

## Author 

Pranjal
