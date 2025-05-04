import { API_HOST, PUBLIC_API_HOST } from '@/shared/config/config';
import { SessionStorage } from '@/shared/lib/session-storage';

export const isBrowser = () => {
  return typeof window !== 'undefined' && typeof window.document !== 'undefined';
}

export const isServer = () => {
  return !isBrowser();
}

export const mutator = async (options: { url: string; method: string; data?: any; headers?: HeadersInit; signal?: AbortSignal }) => {
  const { url, method, data, headers, signal } = options;
  const host = isBrowser() ? PUBLIC_API_HOST : API_HOST;

  // Get session ID from storage
  const sessionId = SessionStorage.getSessionId();

  const requestHeaders: HeadersInit = {
    'Content-Type': 'application/json',
    ...headers
  };

  // Add Authorization header if session ID exists
  if (sessionId) {
    requestHeaders['Authorization'] = `Bearer ${sessionId}`;
  }

  const response = await fetch(API_HOST + url, {
    method,
    headers: requestHeaders,
    body: data ? JSON.stringify(data) : undefined,
    signal
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
};