import React from 'react'
import { MainAppLayout } from '@/widgets/main-app-layout'
import { Card, CardContent, CardHeader, CardTitle } from '@/shared/ui/card'

export const TopPage: React.FC = () => {
  // Placeholder for user data - will be replaced with actual auth context later
  const user = { 
    name: 'John Doe', 
    email: 'john.doe@example.com' 
  }

  return (
    <MainAppLayout>
      <div className="p-4 space-y-4">
        <Card>
          <CardHeader>
            <CardTitle>Welcome</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-lg">
              Hello, {user.name || user.email}!
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