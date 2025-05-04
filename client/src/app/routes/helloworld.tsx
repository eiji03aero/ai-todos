import { createFileRoute } from '@tanstack/react-router'
import { HelloWorldPage } from '@/pages/helloworld/ui/HelloWorldPage'

export const Route = createFileRoute('/helloworld')({
  component: HelloWorldPage
})