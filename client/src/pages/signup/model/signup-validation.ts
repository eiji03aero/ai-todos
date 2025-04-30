import { z } from "zod";

export const signupSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string()
    .min(8, "Password must be at least 8 characters")
    .regex(/[A-Z]/, "Password must contain at least one uppercase letter")
    .regex(/[a-z]/, "Password must contain at least one lowercase letter")
    .regex(/[0-9]/, "Password must contain at least one number")
    .regex(/[!@#$%^&*()]/, "Password must contain at least one special character"),
  passwordConfirmation: z.string()
}).refine((data) => data.password === data.passwordConfirmation, {
  message: "Passwords do not match",
  path: ["passwordConfirmation"]
});

export type SignupFormFields = z.infer<typeof signupSchema>;