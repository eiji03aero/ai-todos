import { $api } from '@/shared/api/client'

export const useSignup = () => {
  return $api.useMutation('post', '/api/auth/signup')
}