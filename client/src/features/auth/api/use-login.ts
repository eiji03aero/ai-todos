import { usePostApiAuthLogin } from '@/shared/api/generated/api';
import { SessionStorage } from '@/shared/lib/session-storage';

export const useLogin = () => {
  return usePostApiAuthLogin({
    mutation: {
      onSuccess: (response) => {
        if (response.sessionId) {
          SessionStorage.setSessionId(response.sessionId);
        }
      },
    },
  });
};