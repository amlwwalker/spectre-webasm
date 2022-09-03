# Avatars

Avatars are user profile pictures.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/avatars.html)


Avatars take four details

1. Size
2. image src
3. Initials
4. Background Colour (displayed when there is no image)

```go
layouts.Columns(
  components.Avatar("avatar-xl", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
  components.Avatar("avatar-lg", "", "AW", "#5755d9"),
  components.Avatar("avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", "AW", "#5755d9"),
  components.Avatar("avatar-xs", "", "AW", "#5755d9"),
),
```
