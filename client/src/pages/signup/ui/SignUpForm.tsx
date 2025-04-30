import { Button } from "@/shared/ui/button"
import { Input } from "@/shared/ui/input"
import { Label } from "@/shared/ui/label"
import { useSignup } from "@/features/auth/api/use-signup"
import { useNavigate } from "@tanstack/react-router"
import { useForm } from "@tanstack/react-form"
import { signupSchema } from "@/pages/signup/model/signup-validation"
import { z } from "zod"

export function SignUpForm() {
  const { mutate, isPending, isError } = useSignup()
  const navigate = useNavigate()

  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
      passwordConfirmation: ''
    },
    onSubmit: async ({ value }) => {
      mutate({
        body: { 
          email: value.email, 
          password: value.password,
          password_confirmation: value.passwordConfirmation
        }
      }, {
        onSuccess: () => {
          navigate({ to: "/" })
        }
      })
    },
    validators: {
      onChange: signupSchema
    }
  })

  return (
    <div className="grid gap-4">
      <form 
        onSubmit={(e) => {
          e.preventDefault()
          form.handleSubmit()
        }}
      >
        <div className="grid gap-2">
          <form.Field
            name="email"
            children={(field) => (
              <div>
                <Label htmlFor={field.name}>Email</Label>
                <Input
                  id={field.name}
                  type="email"
                  placeholder="m@example.com"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                  className={field.state.meta.errors.length ? "border-red-500" : ""}
                />
                {field.state.meta.errors.length > 0 && (
                  <p className="text-red-500 text-sm">
                    {field.state.meta.errors[0]?.message}
                  </p>
                )}
              </div>
            )}
          />

          <form.Field
            name="password"
            children={(field) => (
              <div>
                <Label htmlFor={field.name}>Password</Label>
                <Input
                  id={field.name}
                  type="password"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                  className={field.state.meta.errors.length ? "border-red-500" : ""}
                />
                {field.state.meta.errors.length > 0 && (
                  <p className="text-red-500 text-sm">
                    {field.state.meta.errors[0]?.message}
                  </p>
                )}
              </div>
            )}
          />

          <form.Field
            name="passwordConfirmation"
            children={(field) => (
              <div>
                <Label htmlFor={field.name}>Confirm Password</Label>
                <Input
                  id={field.name}
                  type="password"
                  value={field.state.value}
                  onChange={(e) => field.handleChange(e.target.value)}
                  className={field.state.meta.errors.length ? "border-red-500" : ""}
                />
                {field.state.meta.errors.length > 0 && (
                  <p className="text-red-500 text-sm">
                    {field.state.meta.errors[0]?.message}
                  </p>
                )}
              </div>
            )}
          />

          <Button 
            type="submit" 
            disabled={isPending}
            className="w-full"
          >
            {isPending ? "Signing Up..." : "Sign Up"}
          </Button>

          {isError && (
            <p className="text-red-500 text-sm text-center">
              An error occurred during signup
            </p>
          )}
        </div>
      </form>
    </div>
  )
}