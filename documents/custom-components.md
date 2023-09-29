---
roles: [internal]
group: react
---
# Custom Components

When you want to create your own custom react components, that is very possible.

The way to do this is to create a jsx component. When you want to load it use a custom component and set the data and the component name. It will be rendered automatically.
You can even set custom css against the component in the directory with your component and it will be rendered accordingly.

```go
layouts.Columns(
  components.Avatar("avatar-xl", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
  components.Avatar("avatar-lg", "", "AW", "#5755d9"),
  components.Avatar("avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
  components.Avatar("avatar-xs", "", "AW", "#5755d9"),
),
```


```json
{
  "data": "world"
}
```

* another json block

```json
{
	"data": "content"
}
```

| Syntax      | Description |
| ----------- | ----------- |
| Header      | Title       |
| Paragraph   | Text        |