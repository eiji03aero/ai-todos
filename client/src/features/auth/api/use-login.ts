import { $api } from '@/shared/api';
import { LoginFormData } from '@/pages/login/model/login-validation';

export const useLogin = () => {
  return $api.useMutation('post', '/api/auth/login');
};