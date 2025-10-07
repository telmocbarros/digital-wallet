import apiClient from '../../../api/middleware/auth';

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
  const response = await apiClient.post('/login', credentials);
  return response.data;
}

/**
 * Logout current user
 */
export async function logout(): Promise<void> {
  const response = await apiClient.post('/logout');
  return response.data;
}

/**
 * Check authentication status with server
 */
export async function checkAuthStatus(): Promise<boolean> {
  try {
    const response = await apiClient.get<AuthStatusResponse>('/auth/status');
    return response.data.authenticated;
  } catch (error) {
    console.error('Auth status check failed:', error);
    return false;
  }
}
