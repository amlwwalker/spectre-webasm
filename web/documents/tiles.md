# Tiles

Tiles are repeatable, snippets of information that can contain an icon, a title a subtext or description and a button

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/tabs.html)


```go
app.Div().Class("columns").Body(
  app.Div().Class("column col-9 col-sm-12").Body(
    components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
    components.Tile("The Avengers", "Earth's Mightiest Heroes joined forces to take on threats that were too big for any one hero to tackle...", "Join", "https://picturepan2.github.io/spectre/img/avatar-1.png"),
  ),
),
```
