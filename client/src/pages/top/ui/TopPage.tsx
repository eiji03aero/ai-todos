import React from 'react'
import { MainAppLayout } from '@/widgets/main-app-layout'
import { Card, CardContent, CardHeader, CardTitle } from '@/shared/ui/card'
import { useAppState } from '@/entities/app-state/model/use-app-state'

export const TopPage: React.FC = () => {
  const { appState } = useAppState()

  return (
    <MainAppLayout>
      <div className="p-4 space-y-4">
        <Card>
          <CardHeader>
            <CardTitle>Welcome</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-lg">
              Hello, {appState?.email || 'User'}!
            </p>
            <p className="text-muted-foreground mt-2">
              This is your personal dashboard.
            </p>
          </CardContent>
        </Card>

        <div className="bg-muted p-4 rounded-lg">
          <p className="text-muted-foreground">
            More features are coming soon. Stay tuned!
          </p>
        </div>
      </div>
    </MainAppLayout>
  )
}