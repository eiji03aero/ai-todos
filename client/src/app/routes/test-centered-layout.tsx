import { createFileRoute } from '@tanstack/react-router'
import { TestCenteredLayoutPage } from "@/pages/test-centered-layout/ui/TestCenteredLayoutPage"

export const Route = createFileRoute('/test-centered-layout')({
  component: TestCenteredLayoutPage
})