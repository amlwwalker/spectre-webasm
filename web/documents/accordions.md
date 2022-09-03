# Accordions

Accordions are expanding menus. You can find the original Spectre [documentation here](https://picturepan2.github.io/spectre/components/accordions.html)

The generic accordion structure looks like

```go
components.Accordion(
  // accordion components to add to the accordion
)
```

Then you just add accordion menus to the top level accordion such as

```go
components.Accordion(
   components.AccordionMenu("getting-started", "Getting Started", []string{"Installation", "Custom Version", "Browser-Support"}),
   components.AccordionMenu("elements", "Elements", []string{"Typography", "Tables", "Buttons", "Forms", "Icons.css", "Labels", "Code", "Media"}),
   components.AccordionMenu("layouts", "Layouts", []string{"Flexbox grid", "Responsive", "Hero", "Navbar"}),
   components.AccordionMenu("components", "Components",
      []string{
      "Accordions",
      "Avatars",
      "Badges",
      "Bars",
      "Breadcrumbs",
      "Cards",
      "Chips",
      "Empty States",
      "Menu",
      "Modals",
      "Nav",
      "Pagination",
      "Panels",
      "Popovers",
      "Steps",
      "Tabs",
      "Tiles",
      "Toasts",
      "Tooltips",
      }),
   ),
)
```

An `AccordionMenu` takes three parameters:

1. The ID and path to the page.
2. The title of the accordion menu to be displayed
3. An array of menu items within the accordion.
   1. Note, the menu items names will be lower camel cased for the URL path
