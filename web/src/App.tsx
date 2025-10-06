import './App.css';
import { useEffect, useState } from 'react';
import apiClient, { setAuthState, getAuthState } from './middleware/auth';
function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(getAuthState());
  const [users, setUsers] = useState<{ id: string; email: string }[] | null>(
    null
  );

  useEffect(() => {
    let cancelled = false;
    const checkAuth = async () => {
      const isAuthenticated = await checkAuthStatus();
      if (!cancelled) setIsLoggedIn(isAuthenticated);
    };
    checkAuth();
    return () => {
      cancelled = true;
    };
  }, []);
  async function checkAuthStatus(): Promise<boolean> {
    try {
      const response = await apiClient.get('/auth/status');

      console.log('Auth status response:', response.data);

      const isAuthenticated = response.data.authenticated;
      // Sync client-side state with server reality
      setAuthState(isAuthenticated);

      return isAuthenticated;
    } catch (error) {
      console.log('Not authenticated');
      console.error('Error:', error);
      // Clear auth state on error
      setAuthState(false);
      return false;
    }
  }

  function handleSubmission(e: React.FormEvent) {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;

    apiClient
      .post('/login', { email, password })
      .then((response) => {
        console.log(response.data);
        // Set auth state after successful login
        setAuthState(true);
        setIsLoggedIn(true);
      })
      .catch((error) => {
        console.error('Error:', error);
        // Clear auth state on login failure
        setAuthState(false);
      });
  }

  function fetchUsers() {
    apiClient
      .get('/users')
      .then((response) => {
        console.log(response.data);
        if (response.data.error) {
          // setIsLoggedIn(false);
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

  function handleLogout() {
    apiClient
      .post('/logout')
      .then(() => {
        console.log('Logged out successfully');
        // Clear auth state after successful logout
        setAuthState(false);
        setIsLoggedIn(false);
        setUsers(null);
      })
      .catch((error) => {
        console.error('Logout error:', error);
        // Clear auth state even on error (fail-safe)
        setAuthState(false);
        setIsLoggedIn(false);
        setUsers(null);
      });
  }

  return (
    <div id="content">
      {!isLoggedIn && (
        <form onSubmit={handleSubmission}>
          <label htmlFor="email">Email</label>
          <input type="email" id="email" name="email" autoComplete="on" />
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            autoComplete="on"
          />
          <button type="submit">Login</button>
        </form>
      )}
      {isLoggedIn && (
        <div>
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
        </div>
      )}
    </div>
  );
}

export default App;
