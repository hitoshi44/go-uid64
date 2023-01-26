# UID64

This uid64 package generates 64bit unique identifier, which is Sortable, Clock and Random number-based. Especially, UID can be DB's primary key as long-int(int64) directly.

This package is greatly inspired by [go-scru128](https://github.com/scru128/go-scru128) and [google/uuid](https://github.com/google/uuid). ( But it does NOT follow any RFC specification. )

If you like this package's feature but need more powerful facilities, consider to use [go-scru128](https://github.com/scru128/go-scru128). It is awesome write up about UUID.

## Usage

```go
// Create a NewGenerator with a id that must be in rage 0~3.
g, err := NewGenerator(0)
handle(err)

// Call Gen() for thread-safe generation.
id1, err := g.Gen()
// Call GenDanger() for NOT thread-safe generation.
id2, err := g.GenDanger()

instance := Instance{id1, "test"}

// UID implements sql.Valuer, 
// it can be inserted to sql DB directly, 
// and can be ordered primary key as Long-Int(64bit).
err := db.Save(instance) 
handle(err)

// UID implements sql.Scanner,
// it also can be loaded from DB directly.
```

## Features

Cases this repository is for

 - Sortable, Clock, Random Unique Indentifier
 - Inserted into sql DB as int64 (Long-Int).
 - Thus, SQL Insertion friendly. (it is ordered native integer)
 - Thread-safe, with ID for each genarator. (up to 4)
 - Base36 String. Short, Network friendly string representation.

Cases it's NOT for

 - Will generate more than 128 per a milli-sec.
 - More than 4 distributed generators.
 - Want to follow RFC.

