import { useEffect } from 'react'
import { useNavigate } from '@tanstack/react-router'
import { toast } from 'sonner'
import { usePostApiAuthValidateSession, useGetApiAuthAppState } from '@/shared/api/generated/api'

export function useAuthenticate() {
  const navigate = useNavigate()

  const { mutateAsync } = usePostApiAuthValidateSession()
  const { refetch } = useGetApiAuthAppState({
    query: { enabled: false },
  })

  useEffect(() => {
    async function authenticate() {
      try {
        await mutateAsync()
      } catch (error) {
        toast.error('Session validation failed. Please login again.')
        navigate({ to: '/login' })
        return
      }

      try {
        await refetch()
      } catch (error) {
        toast.error('Failed to fetch app state. Please login again.')
        navigate({ to: '/login' })
      }
    }

    authenticate()
  }, [mutateAsync, refetch, navigate])
}