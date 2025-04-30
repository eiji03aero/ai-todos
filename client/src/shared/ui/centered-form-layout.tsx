import * as React from "react"
import { cn } from "@/shared/lib/shadcn"
import { Card } from "./card"

interface CenteredFormLayoutProps extends React.HTMLAttributes<HTMLDivElement> {
  children: React.ReactNode
}

function CenteredFormLayout({
  children,
  className,
  title,
  ...props
}: CenteredFormLayoutProps) {
  return (
    <div
      className="flex flex-col min-h-screen items-center justify-center p-4"
      {...props}
    >
      {title && (
        <h1 className="text-2xl font-bold mb-4">{title}</h1>
      )}
      <Card
        className={cn(
          "w-full max-w-[480px] mx-auto p-2",
          className
        )}
      >
        {children}
      </Card>
    </div>
  )
}

export { CenteredFormLayout }