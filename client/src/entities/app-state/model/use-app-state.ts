import { create } from 'zustand'
import { useEffect } from 'react'
import { useGetApiAuthAppState } from '@/shared/api/generated/api'
import { GetApiAuthAppState200 } from '@/shared/api/generated/api.schemas'

interface AppStateStore {
  appState: GetApiAuthAppState200 | null
  isAuthenticated: boolean
  setAppState: (state: GetApiAuthAppState200 | null) => void
  setIsAuthenticated: (isAuthenticated: boolean) => void
}

export const useAppStateStore = create<AppStateStore>((set) => ({
  appState: null,
  isAuthenticated: false,
  setAppState: (appState) => set({ appState }),
  setIsAuthenticated: (isAuthenticated) => set({ isAuthenticated })
}))

export const useAppState = () => {
  const { data: appState, isSuccess, refetch } = useGetApiAuthAppState({
    query: {
      enabled: false
    }
  })

  const { setAppState, setIsAuthenticated } = useAppStateStore()

  useEffect(() => {
    if (isSuccess && appState) {
      setAppState(appState)
      setIsAuthenticated(true)
    }
  }, [isSuccess, appState, setAppState, setIsAuthenticated])

  const fetchAppState = () => {
    return refetch()
  }

  return {
    appState: appState ?? null,
    isAuthenticated: isSuccess,
    fetchAppState
  }
}