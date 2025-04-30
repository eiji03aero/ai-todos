import { $api } from '@/shared/api'
import { z } from 'zod'

const emailSchema = z.string().email('Invalid email format')

export function useCheckEmailUniqueness(email: string) {
  return $api.useQuery('get', '/api/auth/check-email', {
    params: {
      query: { email }
    },
    query: () => {
      // Validate email format first
      try {
        emailSchema.parse(email)
      } catch (validationError) {
        throw new Error('Invalid email format')
      }
    },
    enabled: !!email, // Only run query if email is not empty
  })
}