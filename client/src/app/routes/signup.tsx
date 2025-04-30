import { createFileRoute } from '@tanstack/react-router'
import { SignUpPage } from '@/pages/signup/ui/SignUpPage'

export const Route = createFileRoute('/signup')({
  component: SignUpPage,
})
