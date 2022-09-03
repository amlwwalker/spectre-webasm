# Badges

Badges are often used to inform the user of a value or changing score.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/badges.html)

Badges could potentially listen for changes automatically

Badges can be used in one of three pre-defined ways

```go
components.NotificationBadge("notification", 7),
components.ButtonBadge("button", 7, nil),
components.FigureBadge("https://picturepan2.github.io/spectre/img/avatar-3.png", 7, nil),
```

each of which displays a different type of badge.

`FigureBadge` and `ButtonBadge` can take an onClick mechanism to convert the avatar into a clickable.

## OnClick

As with any other onClick, the prototype of the function to handle clicks is

```go
func(ctx app.Context, e app.Event)
```
