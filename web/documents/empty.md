# Empty States

An empty state component can include icons, messages (title and subtitle messages) and action buttons or any combination of those elements.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/empty.html)

```go
components.EmptyState("icon-people", "An Empty State", "An example empty state", "btn-primary", "does nothing", nil),
components.EmptyState("icon-3x icon-mail", "You have no new messages", "Click the button to start a conversation", "btn-primary", "Send a message", nil),
```

