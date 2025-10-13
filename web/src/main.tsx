import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import './api/middleware/auth.ts';
import App from './App.tsx';

import { createBrowserRouter, RouterProvider } from 'react-router';
import MainLayout from './components/templates/MainLayout.tsx';
import Dashboard from './pages/Dashboard.tsx';
import { TestPage } from './pages/TestPage.tsx';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    element: <MainLayout />,
    children: [
      {
        path: '/dashboard',
        element: <Dashboard />,
      },
      {
        path: '/test',
        element: <TestPage />,
      },
    ],
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
