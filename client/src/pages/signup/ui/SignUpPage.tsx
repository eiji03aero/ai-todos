import { CenteredFormLayout } from "@/shared/ui/centered-form-layout"
import { SignUpForm } from "@/pages/signup/ui/SignUpForm"
import { toast } from 'sonner'

export function SignUpPage() {
  return (
    <CenteredFormLayout title="Signup for AI todos">
      <div className="p-6">
        <SignUpForm />
      </div>
    </CenteredFormLayout>
  )
}