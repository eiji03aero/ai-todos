import { createFileRoute } from '@tanstack/react-router'
import { TopPage } from '@/pages/top'

export const Route = createFileRoute('/')({
  component: TopPage,
})
