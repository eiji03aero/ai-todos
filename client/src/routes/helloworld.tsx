import { createFileRoute } from '@tanstack/react-router'
import { HelloWorldPage } from '../pages/helloworld/ui/HelloWorldPage'

function HelloWorldRoute() {
  return <HelloWorldPage />
}

export const Route = createFileRoute('/helloworld')({
  component: HelloWorldRoute,
})