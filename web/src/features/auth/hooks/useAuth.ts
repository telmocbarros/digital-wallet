import { useState, useEffect } from 'react';
import { login as loginService, logout as logoutService, checkAuthStatus } from '../api/authService';
import { setAuthState as setAuthStateStorage, getAuthState as getAuthStateStorage } from '../../../api/middleware/auth';

export function useAuth() {
  const [isLoggedIn, setIsLoggedIn] = useState(getAuthStateStorage());
  const [isLoading, setIsLoading] = useState(true);

  // Check auth status on mount
  useEffect(() => {
    let cancelled = false;

    const verifyAuth = async () => {
      const isAuthenticated = await checkAuthStatus();
      
      if (!cancelled) {
        // Sync client-side state with server reality
        setAuthStateStorage(isAuthenticated);
        setIsLoggedIn(isAuthenticated);
        setIsLoading(false);
      }
    };

    verifyAuth();

    return () => {
      cancelled = true;
    };
  }, []);

  const login = async (email: string, password: string) => {
    try {
      await loginService({ email, password });
      setAuthStateStorage(true);
      setIsLoggedIn(true);
    } catch (error) {
      console.error('Login failed:', error);
      setAuthStateStorage(false);
      throw error;
    }
  };

  const logout = async () => {
    try {
      await logoutService();
    } catch (error) {
      console.error('Logout error:', error);
    } finally {
      // Clear auth state even on error (fail-safe)
      setAuthStateStorage(false);
      setIsLoggedIn(false);
    }
  };

  return {
    isLoggedIn,
    isLoading,
    login,
    logout,
  };
}
