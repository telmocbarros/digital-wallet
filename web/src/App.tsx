import './App.css';
import { useEffect, useState } from 'react';
function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

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
      const response = await fetch(`${import.meta.env.VITE_API_ENDPOINT}/`, {
        method: 'GET',
        credentials: 'include', // Include cookies in the request
      });

      if (!response.ok) {
        console.log('Not authenticated');
        return false;
      }

      const data = await response.json();
      console.log('User is authenticated: ', data);

      return true;
    } catch (error) {
      console.error('Error:', error);
      return false;
    }
  }

  function handleSubmission(e: React.FormEvent) {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;

    fetch(`${import.meta.env.VITE_API_ENDPOINT}/login`, {
      method: 'POST',
      credentials: 'include', // Include cookies in the request
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setIsLoggedIn(true);
      })
      .catch((error) => {
        console.error('Error:', error);
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
      {isLoggedIn && <p>Welcome back!</p>}
    </div>
  );
}

export default App;
