import './App.css';
import { useState } from 'react';
import apiClient from './api/middleware/auth';
import { useAuth } from './features/auth/hooks/useAuth';
import LoginForm from './features/auth/components/LoginForm';
import AuthLayout from './components/templates/AuthLayout';
import MainLayout from './components/templates/MainLayout';

function App() {
  const { isLoggedIn, login, logout } = useAuth();
  const [users, setUsers] = useState<{ id: string; email: string }[] | null>(
    null
  );

  async function handleLogin(email: string, password: string) {
    try {
      await login(email, password);
      console.log('Login successful');
    } catch (error) {
      console.error('Login failed:', error);
    }
  }

  function fetchUsers() {
    apiClient
      .get('/users')
      .then((response) => {
        console.log(response.data);
        if (response.data.error) {
          setUsers(null);
          return;
        }
        setUsers(response.data);
      })
      .catch((error) => {
        console.error('Error:', error);
        setUsers(null);
      });
  }

  async function handleLogout() {
    await logout();
    setUsers(null);
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
          <button onClick={fetchUsers}>Fetch Users</button>
          <button onClick={handleLogout}>Logout</button>
          <ol>
            {users &&
              users.map((user: { id: string; email: string }) => (
                <li key={user.id}>
                  {user.id} - {user.email}
                </li>
              ))}
          </ol>
        </MainLayout>
      )}
    </>
  );
}

export default App;
