import React, { useState } from "react"
import { 
  CenteredFormLayout, 
  Card, 
  CardContent, 
  CardHeader, 
  CardTitle,
  Input,
  Label,
  Button 
} from "@/shared/ui"

export function TestCenteredLayoutPage() {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: value
    }))
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    alert(JSON.stringify(formData, null, 2))
  }

  return (
    <CenteredFormLayout>
      <CardHeader>
        <CardTitle>Test Form</CardTitle>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="username">Username</Label>
            <Input 
              id="username"
              name="username"
              value={formData.username}
              onChange={handleChange}
              placeholder="Enter username"
              required
            />
          </div>
          
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input 
              id="email"
              name="email"
              type="email"
              value={formData.email}
              onChange={handleChange}
              placeholder="Enter email"
              required
            />
          </div>
          
          <Button type="submit" className="w-full">Submit</Button>
        </form>
      </CardContent>
    </CenteredFormLayout>
  )
}