import { z } from 'zod';
import { useGetApiAuthCheckEmail } from '@/shared/api/generated/api';

const emailSchema = z.string().email('Invalid email format');

export function useCheckEmailUniqueness(email: string) {
  return useGetApiAuthCheckEmail({ email }, {
    query: {
      enabled: !!email,
    },
  });
}