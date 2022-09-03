# Tabs

Tabs allow organising information, or grouping information based on content or context.

You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/tabs.html)

```go
app.Div().Class("column col-12").Body(
  components.Tabs(
    components.Tab("Profile", h.activeTab == "profile-tab", func(ctx app.Context, e app.Event) {
      h.activeTab = "profile-tab"
      h.Update()
    }),
    components.Tab("Files", h.activeTab == "files-tab", func(ctx app.Context, e app.Event) {
      h.activeTab = "files-tab"
      h.Update()
    }),
    components.Tab("Tasks", h.activeTab == "tasks-tab", func(ctx app.Context, e app.Event) {
      h.activeTab = "tasks-tab"
      h.Update()
    }),
  ),
  app.If(h.activeTab == "profile-tab",
    components.Card("Profile", "Profile of user", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
      app.Span().Body(
        app.Span().Text("Profile of a user to deal with"),
      ),
      app.Span().Body(
        app.Button().Class("btn btn-primary").Text("Card Button"),
      ),
    ),
  ),
  app.If(h.activeTab == "files-tab",
    components.Card("Files", "Files that need handling", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
      app.Span().Body(
        app.Span().Text("Files that the user hasn't processed yet"),
      ),
      app.Span().Body(
        app.Button().Class("btn btn-primary").Text("Card Button"),
      ),
    ),
  ),
  app.If(h.activeTab == "tasks-tab",
    components.Card("Tasks", "Tasks that need doing", "https://picturepan2.github.io/spectre/img/osx-yosemite.jpg",
      app.Span().Body(
        app.Span().Text("Tasks the user is yet to deal with"),
      ),
      app.Span().Body(
        app.Button().Class("btn btn-primary").Text("Card Button"),
      ),
    ),
  ),
),
```
