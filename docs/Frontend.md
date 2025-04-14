# Frontend Development Guidelines

## Technology Stack

### Core Technologies
- **Package manager**: npm
- **Language**: TypeScript
- **UI Framework**: React
- **React Framework**: Tanstack Start
- **UI Library**: shadcn/ui
- **Styling**: 
  - Tailwind CSS
  - CSS Modules
- **Form Management**: Tanstack Form
- **HTTP Communication**: 
  - Axios (HTTP Client)
  - Tanstack Query (State Management for HTTP Requests)

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
- segments
  - segments group your code by its purpose.
  - any name can be given to a segment, but usually we use one of these:
    - ui — everything related to UI display: UI components, date formatters, styles, etc.
    - api — backend interactions: request functions, data types, mappers, etc.
    - model — the data model: schemas, interfaces, stores, and business logic.
    - lib — library code that other modules on this slice need.
    - config — configuration files and feature flags.

## Folder structure
Please refer to following for understanding folder structure.
Let's say we have `auth` as a slice, `ui` and `model` as segments.

- client/src/
  - app/: FSD's app layer
    - none
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
