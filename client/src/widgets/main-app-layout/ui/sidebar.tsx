import React from 'react'
import { Button } from '@/shared/ui/button'
import { 
  NavigationMenu, 
  NavigationMenuList, 
  NavigationMenuItem, 
  NavigationMenuLink 
} from '@/shared/ui/navigation-menu'

export const Sidebar: React.FC = () => {
  return (
    <aside className="w-64 bg-muted p-4 border-r">
      <NavigationMenu orientation="vertical">
        <NavigationMenuList>
          <NavigationMenuItem>
            <NavigationMenuLink asChild>
              <Button variant="ghost" className="w-full justify-start">
                Dashboard
              </Button>
            </NavigationMenuLink>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuLink asChild>
              <Button variant="ghost" className="w-full justify-start">
                Tasks
              </Button>
            </NavigationMenuLink>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuLink asChild>
              <Button variant="ghost" className="w-full justify-start">
                Settings
              </Button>
            </NavigationMenuLink>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </aside>
  )
}