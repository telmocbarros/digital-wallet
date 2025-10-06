import axios from 'axios';

const PUBLIC_API_ENDPOINTS = ['/login', '/refresh', '/register', '/auth/status'];
// Middleware to handle refresh token logic

export const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_ENDPOINT,
  withCredentials: true, // Include cookies in requests
});

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
        // Retry the original request
        return apiClient(originalRequest);
      } catch (refreshError) {
        // If refresh fails, redirect to login or handle accordingly
        console.error('Token refresh failed:', refreshError);
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default apiClient;
