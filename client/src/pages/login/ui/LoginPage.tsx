import { CenteredFormLayout } from '@/shared/ui/centered-form-layout';
import { LoginForm } from './LoginForm';

export const LoginPage = () => {
  return (
    <CenteredFormLayout title="AI Todos Login">
      <LoginForm />
    </CenteredFormLayout>
  );
};