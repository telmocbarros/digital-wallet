import axios from 'axios';

const PUBLIC_API_ENDPOINTS = ['/login', '/refresh', '/register', '/auth/status'];
const AUTH_FLAG_KEY = 'isAuthenticated';

// Helper functions for auth state management
export const setAuthState = (isAuthenticated: boolean) => {
  if (isAuthenticated) {
    localStorage.setItem(AUTH_FLAG_KEY, 'true');
  } else {
    localStorage.removeItem(AUTH_FLAG_KEY);
  }
};

export const getAuthState = (): boolean => {
  return localStorage.getItem(AUTH_FLAG_KEY) === 'true';
};

export const clearAuthState = () => {
  localStorage.removeItem(AUTH_FLAG_KEY);
};

// Middleware to handle refresh token logic
export const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_ENDPOINT,
  withCredentials: true, // Include cookies in requests
});

// Request interceptor to skip protected requests when not authenticated
apiClient.interceptors.request.use(
  (config) => {
    // Check if this is a public endpoint
    const isPublicEndpoint = PUBLIC_API_ENDPOINTS.some((endpoint) =>
      config.url?.includes(endpoint)
    );

    // If it's not a public endpoint and we know the user isn't authenticated,
    // reject the request early to avoid unnecessary network calls
    if (!isPublicEndpoint && !getAuthState()) {
      console.log('Skipping protected request - user not authenticated');
      return Promise.reject(new axios.Cancel('User not authenticated'));
    }

    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor to handle 401 errors and refresh token logic
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // Check if the failed request was to a public endpoint
    const isPublicEndpoint = PUBLIC_API_ENDPOINTS.some((endpoint) =>
      originalRequest.url?.includes(endpoint)
    );

    // Skip retry logic if:
    // 1. The endpoint is public (expected to return 401 when not authenticated)
    // 2. We've already tried retrying once
    if (
      error.response &&
      error.response.status === 401 &&
      !originalRequest._retry &&
      !isPublicEndpoint
    ) {
      originalRequest._retry = true;
      try {
        // Attempt to refresh the token
        await apiClient.post('/refresh');
        // Token refresh succeeded, retry the original request
        return apiClient(originalRequest);
      } catch (refreshError) {
        // If refresh fails, the session is truly dead
        console.error('Token refresh failed:', refreshError);
        // Clear auth state - user needs to log in again
        clearAuthState();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default apiClient;
