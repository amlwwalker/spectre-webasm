---
roles: [maintainer, user]
group: components
---
# Chips

Chips are complex entities in small blocks.

```go
layouts.FlexBox("80%",
  components.Chip("Crime", "", "", false),
  components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", false),
  components.Chip("Crime", "avatar-sm", "https://picturepan2.github.io/spectre/img/avatar-1.png", true),
),
```
