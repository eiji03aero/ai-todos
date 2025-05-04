import React, { ReactNode } from 'react'
import { Sidebar } from '@/widgets/main-app-layout/ui/sidebar'
import { useAuthenticate } from '@/features/auth/model/useAuthenticate'

interface MainAppLayoutProps {
  children: ReactNode
}

export const MainAppLayout: React.FC<MainAppLayoutProps> = ({ children }) => {
  useAuthenticate()

  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <main className="flex-1 p-4 bg-background">
        {children}
      </main>
    </div>
  )
}