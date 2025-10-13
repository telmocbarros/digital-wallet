import apiClient from '../../../api/middleware/auth';

/** Dummy For Development - START */
import { dummy_users } from '../../../database/dummy_data';
/** Dummy For Development - END */

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface AuthStatusResponse {
  authenticated: boolean;
}

/**
 * Login user with email and password
 */
export async function login(credentials: LoginCredentials): Promise<void> {
  // NOTE: Dummy For Development - START
  const user = dummy_users.find(
    (u) => u.email === credentials.email && u.password === credentials.password
  );
  if (!user) {
    return Promise.reject(new Error('Invalid email or password'));
  }
  return Promise.resolve();
  // NOTE: Dummy For Development - END

  // const response = await apiClient.post('/login', credentials);

  // return response.data;
}

/**
 * Logout current user
 */
export async function logout(): Promise<void> {
  // NOTE: Dummy For Development - START
  return Promise.resolve();
  // NOTE: Dummy For Development - END

  //
  // const response = await apiClient.post('/logout');
  // return response.data;
}

/**
 * Check authentication status with server
 */
export async function checkAuthStatus(): Promise<boolean> {
  // NOTE: Dummy For Development - START
  return Promise.resolve(true);
  // NOTE: Dummy For Development - END
  // try {
  //   const response = await apiClient.get<AuthStatusResponse>('/auth/status');
  //   return response.data.authenticated;
  // } catch (error) {
  //   console.error('Auth status check failed:', error);
  //   return false;
  // }
}
