# Frontend Development Guidelines

## Technology Stack

### Core Technologies
- **Package manager**: npm
- **Language**: typescript
- **UI Framework**: react
- **React Framework**: @tanstack/start
- **UI Library**: shadcn/ui
- **Styling**: 
  - tailwind
  - CSS Modules
- **Form Management**: @tanstack/react-form
- **Validation**: zod
- **HTTP Communication**: 
  - openapi-fetch
  - @tanstack/react-query (State Management for HTTP Requests)

## Project Structure: Feature Sliced Design

The project follows the Feature Sliced Design (FSD) architectural methodology, which provides a clear, scalable, and maintainable folder structure. The FSD structure is implemented under the `client/src/` directory.

Consists of these three pillaars:
- layers
  - consists of 7 standardized divisions:
    - app
      - everything that makes the app run — routing, entrypoints, global styles, providers. 
      - directly has segments, no slices
    - processes
      - deprecated, don't use this
    - pages
      - full pages or large parts of a page in nested routing.
    - widgets
      - large self-contained chunks of functionality or UI, usually delivering an entire use case.
    - features
      - reused implementations of entire product features, i.e. actions that bring business value to the user.
    - entities
      - business entities that the project works with, like user or product.
    - shared
      - reusable functionality, especially when it's detached from the specifics of the project/business, though not necessarily.
      - directly has segments, no slices
  - modules on one layer can only know about and import from modules from the layers strictly below.
  - basically layers consists of slices, then each of those slices consists of segments
    - exception is that app and shared layers directly have segments
- slices
  - partitions the code by business domain
  - slices cannot use other slices on the same layer, and that helps with high cohesion and low coupling.
  - a slice should have a index.ts file that exports implementations of the child segments that should be pubic to upper layers
- segments
  - segments group your code by its purpose.
  - any name can be given to a segment, but usually we use one of these:
    - ui — everything related to UI display: UI components, date formatters, styles, etc.
    - api — backend interactions: request functions, data types, mappers, etc.
    - model — the data model: schemas, interfaces, stores, and business logic.
    - lib — library code that other modules on this slice need.
    - config — configuration files and feature flags.
  - for the segments under shared layer, each segment should have a index.ts file that exports implementations that should be public to upper layers

## Folder structure
Please refer to following for understanding folder structure.
Let's say we have `auth` as a slice, `ui` and `model` as segments.

- client/src/
  - app/: FSD's app layer
    - routes
      - contains route files based on tanstack/start library
      - using file based routes
  - pages/: FSD's pages layer
    - auth/
      - ui/
        - LoginPage.tsx
  - widgets/: FSD's widgets layer
    - auth/
      - ui/
        - none
      - model/
        - none
  - features/: FSD's features layer
    - auth/
      - ui/
        - LoginForm.tsx
      - model/
        - useLogin.ts
  - entities/: FSD's entities layer
    - auth/
      - model/
        - User.ts
  - shared/: FSD's shared layer
    - ui
      - Button.tsx

## Development Principles

- **Isolation**: Each feature should be self-contained and minimally coupled
- **Scalability**: Easy to add new features without modifying existing code
- **Maintainability**: Clear separation of concerns
- **Reusability**: Promote component and logic reuse

## Best Practices

- Use TypeScript for type safety
- Leverage Tanstack Query for efficient data fetching and caching
- Utilize Tailwind CSS for rapid styling
- Implement CSS Modules for component-specific styles
- Use shadcn/ui for consistent and accessible UI components
- Follow React best practices and functional component patterns

## Impelmentation
### How to define a route
- Create a corresponding page implementation in client/src/pages/[slice]/ui/[component page name].tsx
- Register route by following step
  - create a file under client/routes/[path-to-component].tsx
    - put the page path to the path-to-component part
  - register the route implementing a code like following
    ```
    import { createFileRoute } from '@tanstack/react-router'
    import { SamplePage } from 'path-to-page-component';

    export const Route = createFileRoute('/path-path')({
      component: SamplePage,
    })
    ```
  - if you get a type error for the first argument of createFileRoute, run npm run dev in docker-compose's client container, so that it would generate type definition file for routing

### File name and export
- When implementing one of the following:
  - react component
  - react custom hook
- Follow these rules
  - Name the file after the target code's name(component, hook ...)
  - export the target code with named export

### Designing implementation
- Instead of directly putting implementations in pages layer, try to utilize widgets, features, entities, shared layers as much as possible
- If you have doubts to which layer to put, ask the user for opinion

### React implementation
- Try to extract component / hooks / functions if you think they have reusability.
  - Extract them to the other layers than pages, like widgets, features, entities, shared
- Try to extract logics to react custom hook from component on purpose basis.

### Layout to apply
We have following layout implementation for this app:

- Centered form
  - Has a card centered in the page that can have an arbitrary content inside of it
  - use this for:
    - non authenticated page which needs a form, like signup page
- Main app layout
  - This is the layout we apply to the authenticated application pages
  - Has a side navigation on left
  - to the right of side navigation is the main content area of the page
  - use this for:
    - any of the authenticated application pages

### Change rule
- do not attempt to modify *.gen.ts file as they are auto genearted file

### To implement a page
- When implementing a page, follow below rules:
  - create actual implementation on page layer with a dedicated slice
  - create a route in app/routes folder while sticking with tanstack router's file based router

### import rules
- when import from segments or slices, make sure to import from the index.ts file so that we will stick with the ones that are specified as public apis
- when writing import path, please use path alias, do not use relative path.
  - OK (using path alias): @/shared/ui
  - NG (using relative path): ../../shared/ui

### Base UI components
- Please consider to use shadcn components for basic UI necessities
- They are located in client/src/shared/ui

### To implement a form
- Please use @tanstack/react-form and zod
- pass validation schema created by zod to useForm's `validators.onChange` property

### Implement http request
- Use openapi-fetch. wrapped implementation `$api` is exported from src/shared/api
- When implement a http request, create a custom hook that uses openapi-fetch's react-query integrations like useQuery, useMutation provided by the wrapped implementation `$api`
  - You don't need to pass type arguments to these integrations useQuery, useMutation
  - please check following documents for how to use respective integrations
    - $api.useQuery: https://openapi-ts.dev/openapi-react-query/use-query
      - Reference:
        ```
          const { data, error, isLoading } = $api.useQuery("get", "/users/{user_id}", {
            params: {
              path: { user_id: 5 },
            },
          });
        ```
    - $api.useMutation: https://openapi-ts.dev/openapi-react-query/use-mutation
      - please use lowercase for the method
      - Reference
        ```
        // argument reference
        const mutation = $api.useMutation(method, path, queryOptions, queryClient);
        // example usage
        const mutation = $api.useMutation("patch", "/users");
        ```
- When using react-query integrations like useQuery, useMutation, please make sure to inspect and understand the type signature of the provided hook so that you won't pass wrong arguments
- Try to define the custom hook in features layer under proper slice

## Running npm commands
- Please run them inside client container of docker compose