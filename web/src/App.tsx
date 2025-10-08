import './App.css';
import { useEffect } from 'react';
import { useAuth } from './features/auth/hooks/useAuth';
import { useTheme } from './hooks/useTheme';
import LoginForm from './features/auth/components/LoginForm';
import AuthLayout from './components/templates/AuthLayout';
import MainLayout from './components/templates/MainLayout';

function App() {
  const { /*isLoggedIn,*/ login /*, logout*/ } = useAuth();
  const isLoggedIn = true; // For testing purposes
  const { theme } = useTheme();

  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
  }, [theme]);

  async function handleLogin(email: string, password: string) {
    try {
      await login(email, password);
      console.log('Login successful');
    } catch (error) {
      console.error('Login failed:', error);
    }
  }

  return (
    <>
      {!isLoggedIn && (
        <AuthLayout>
          <LoginForm onSubmit={handleLogin} />
        </AuthLayout>
      )}
      {isLoggedIn && (
        <MainLayout>
          <h1>Do Something</h1>
          <button
            onClick={() => {
              alert('Clicked!');
            }}
          >
            Click Me
          </button>
        </MainLayout>
      )}
    </>
  );
}

export default App;
