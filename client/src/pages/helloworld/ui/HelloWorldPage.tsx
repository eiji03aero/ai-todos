import React from 'react'
import { HelloWorlder } from '@/widgets/helloworld/ui/HelloWorlder'

export const HelloWorldPage: React.FC = () => {
  return (
    <div className="container mx-auto p-4">
      <div className="flex flex-col items-center justify-center space-y-6">
        <h1 className="text-4xl font-bold text-blue-600">Hello, World!</h1>
        <HelloWorlder />
      </div>
    </div>
  )
}