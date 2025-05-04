const SESSION_KEY = 'session_id'

export const SessionStorage = {
  setSessionId: (sessionId: string) => {
    localStorage.setItem(SESSION_KEY, sessionId)
  },
  
  getSessionId: () => {
    return localStorage.getItem(SESSION_KEY)
  },
  
  removeSessionId: () => {
    localStorage.removeItem(SESSION_KEY)
  },
  
  hasValidSession: () => {
    return !!localStorage.getItem(SESSION_KEY)
  }
}