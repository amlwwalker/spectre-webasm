---
roles: [internal]
group: elements
---
# Tables

Tables are straight forward to create and will fit into any layout structure

Tables take advantage of Go generics so anything can be placed in a table. The struct fields will be selected uniquely as the table column names

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/elements/tables.html)

Constructing a table

```go
t := []any{
  tmp{
    A: "a",
    B: "b",
    C: "c1",
  },
  tmp2{
    C: "c2",
    D: "d",
    E: "e",
  },
}
components.Table(t, "table-striped")
```


The order of the elements defined in the structs will affect the order in which they appear as columns in the table. The first time they appear in a struct will affect its position in the table.
