import { usePostApiAuthSignup } from '@/shared/api/generated/api';
import { SessionStorage } from '@/shared/lib/session-storage';

export const useSignup = () => {
  return usePostApiAuthSignup({
    mutation: {
      onSuccess: (response) => {
        if (response.sessionId) {
          SessionStorage.setSessionId(response.sessionId);
        }
      },
    },
  });
};